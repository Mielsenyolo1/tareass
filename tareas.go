package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type listaTareas struct {
	tareas []Tarea
}

// NuevaTarea crea una nueva tarea con el ID proporcionado.
func (l *listaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *listaTareas) completado(idex int) {

	l.tareas[idex].completado = true
}

func (l *listaTareas) editarTarea(idex int, t Tarea) {
	l.tareas[idex] = t

}
func (l *listaTareas) eliminarTarea(idex int) {
	l.tareas = append(l.tareas[:idex], l.tareas[idex+1:]...)
}

func main() {

	//instancia de lista de tareas
	lista := listaTareas{}

	//instacia para entrar tareas
	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Agregar tarea")
		fmt.Println("2. Editar tarea")
		fmt.Println("3. Eliminar tarea")
		fmt.Println("4. Marcar tarea como completada")
		fmt.Println("5. salir")
		fmt.Println("elija una opción:")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Println("Ingrese el nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada exitosamente.")
		case 2:
			var idex int
			var t Tarea
			fmt.Println("Ingrese el índice de la tarea a editar:")
			fmt.Scanln(&idex)
			fmt.Println("Ingrese el nuevo nombre de la tarea:")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Println("Ingrese la nueva descripción de la tarea:")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(idex, t)
			fmt.Println("Tarea editada exitosamente.")

		case 3:
			var index int
			fmt.Println("Ingrese el índice de la tarea a eliminar:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada exitosamente.")

		case 4:
			var index int
			fmt.Println("Ingrese el índice de la tarea a marcar como completada:")
			fmt.Scanln(&index)
			lista.completado(index)
			fmt.Println("Tarea marcada como completada.")

		case 5:
			fmt.Println("Saliendo del programa.")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}

		//listar tareas
		fmt.Println("Bienvenido a la lista de tareas")
		fmt.Println("==============================================================================")
		for i, t := range lista.tareas {

			fmt.Printf("%d. %s - %s - completado: %t\n", i, t.nombre, t.descripcion, t.completado)

		}
		fmt.Println("==============================================================================")
	}

}
