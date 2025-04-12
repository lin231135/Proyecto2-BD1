package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
)

const (
	defaultDSN  = "postgres://postgres:12345@localhost:5432/postgres?sslmode=disable"
	proyectoDSN = "postgres://postgres:12345@localhost:5432/Proyecto_2?sslmode=disable" //CAMBIAR EL USUARIO Y CONTRASEÑA SEGUN SU CONFIGURACION
)

func main() {
	fmt.Println("¿Deseas reiniciar la base de datos con ddl.sql y data.sql? (s/n)")
	var respuesta string
	fmt.Scanln(&respuesta)
	if strings.ToLower(respuesta) == "s" {
		if err := inicializarBaseDeDatos(); err != nil {
			fmt.Println("Error inicializando la base de datos:", err)
			return
		}

		if err := ejecutarScriptSQL("db/ddl.sql"); err != nil {
			fmt.Println("Error ejecutando ddl.sql:", err)
			return
		}
		if err := ejecutarScriptSQL("db/data.sql"); err != nil {
			fmt.Println("Error ejecutando data.sql:", err)
			return
		}
		fmt.Println("Base de datos reiniciada correctamente.\n")
	}

	var opcionUsuarios, opcionAislamiento, usuarios int
	var isoLevel pgx.TxIsoLevel
	var aislamientoStr string

	fmt.Println("--- Simulación de Reservas Concurrentes ----")
	fmt.Println("Seleccione el número de usuarios concurrentes:")
	fmt.Println("1. 5 usuarios")
	fmt.Println("2. 10 usuarios")
	fmt.Println("3. 20 usuarios")
	fmt.Println("4. 30 usuarios")
	fmt.Print("Opción: ")
	fmt.Scanln(&opcionUsuarios)

	switch opcionUsuarios {
	case 1:
		usuarios = 5
	case 2:
		usuarios = 10
	case 3:
		usuarios = 20
	case 4:
		usuarios = 30
	default:
		fmt.Println("Opción inválida, se usarán 5 usuarios por defecto.")
		usuarios = 5
	}

	fmt.Println("\nSeleccione el nivel de aislamiento:")
	fmt.Println("1. READ COMMITTED")
	fmt.Println("2. REPEATABLE READ")
	fmt.Println("3. SERIALIZABLE")
	fmt.Print("Opción: ")
	fmt.Scanln(&opcionAislamiento)

	switch opcionAislamiento {
	case 1:
		isoLevel, aislamientoStr = pgx.ReadCommitted, "READ COMMITTED"
	case 2:
		isoLevel, aislamientoStr = pgx.RepeatableRead, "REPEATABLE READ"
	case 3:
		isoLevel, aislamientoStr = pgx.Serializable, "SERIALIZABLE"
	default:
		fmt.Println("Opción inválida, se usará SERIALIZABLE por defecto.")
		isoLevel, aislamientoStr = pgx.Serializable, "SERIALIZABLE"
	}

	var wg sync.WaitGroup
	wg.Add(usuarios)

	start := time.Now()
	var successCount, failCount int
	var mutex sync.Mutex

	for i := 1; i <= usuarios; i++ {
		go func(usuarioID int) {
			defer wg.Done()
			if err := intentarReserva(usuarioID, isoLevel); err != nil {
				mutex.Lock()
				failCount++
				mutex.Unlock()
				fmt.Printf("Usuario %d: fallo -> %v\n", usuarioID, err)
			} else {
				mutex.Lock()
				successCount++
				mutex.Unlock()
				fmt.Printf("Usuario %d: éxito\n", usuarioID)
			}
		}(i)
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Println("\n---------------- Resultados ----------------")
	fmt.Printf("Usuarios Concurrentes: %d\n", usuarios)
	fmt.Printf("Nivel de Aislamiento: %s\n", aislamientoStr)
	fmt.Printf("Reservas Exitosas: %d\n", successCount)
	fmt.Printf("Reservas Fallidas: %d\n", failCount)
	fmt.Printf("Tiempo Promedio: %d ms\n", elapsed.Milliseconds())
}

func inicializarBaseDeDatos() error {
	conn, err := pgx.Connect(context.Background(), defaultDSN)
	if err != nil {
		return fmt.Errorf("error conectando a postgres: %v", err)
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), `DROP DATABASE IF EXISTS "Proyecto_2"`)
	if err != nil {
		return fmt.Errorf("error al eliminar la base de datos: %v", err)
	}

	_, err = conn.Exec(context.Background(), `CREATE DATABASE "Proyecto_2"`)
	if err != nil {
		return fmt.Errorf("error al crear la base de datos: %v", err)
	}

	return nil
}

func intentarReserva(usuarioID int, iso pgx.TxIsoLevel) error {
	conn, err := pgx.Connect(context.Background(), proyectoDSN)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{IsoLevel: iso})
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	var existe bool
	err = tx.QueryRow(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM reserva WHERE id_asiento = $1)", 2).Scan(&existe)
	if err != nil {
		return err
	}
	if existe {
		return fmt.Errorf("asiento ya reservado")
	}

	_, err = tx.Exec(context.Background(),
		"INSERT INTO reserva (id_usuario, id_asiento, fecha_reserva) VALUES ($1, $2, NOW())",
		usuarioID, 2)
	if err != nil {
		return err
	}

	return tx.Commit(context.Background())
}

func ejecutarScriptSQL(filename string) error {
	sqlBytes, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("no se pudo leer %s: %v", filename, err)
	}
	sql := string(sqlBytes)

	conn, err := pgx.Connect(context.Background(), proyectoDSN)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return fmt.Errorf("error al ejecutar script SQL: %v", err)
	}

	return nil
}
