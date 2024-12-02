package files

import (
	"encoding/json"
	"fmt"
	"os"

	"trucode.com/challenge-w5/users"
)

type InfoUsers struct {
	Users []users.User `json:"users"`
}

// Lee los datos del archivo JSON
func ReadData(filename string) (InfoUsers, error) {
	var dataUsers InfoUsers

	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return dataUsers, nil
	}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return dataUsers, err
	}

	err = json.Unmarshal(bytes, &dataUsers)
	return dataUsers, err
}

// Atualiza archivo json
func WriteData(filename string, dataUsers InfoUsers) error {
	bytes, err := json.MarshalIndent(dataUsers, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, bytes, 0644)
	return err
}

// Crear archivo nuevo si no existiera ningun usuario
func InitializeData(filename string, dataUsers *InfoUsers, userName string) error {
	if len(dataUsers.Users) != 0 {
		return nil
	}

	fmt.Println("No se encontraron usuarios. Inicializando estructura...")
	*dataUsers = InfoUsers{
		Users: []users.User{
			users.NewUser(userName),
		},
	}

	if err := WriteData(filename, *dataUsers); err != nil {
		return fmt.Errorf("error al crear el archivo: %w", err)
	}

	fmt.Println("Archivo JSON creado con la estructura básica.")
	return nil
}

// añade usuario con su estructura al json
func AddUserInJson(filename string, userName string) error {
	dataUsers, err := ReadData(filename)
	if err != nil {
		return fmt.Errorf("error al leer el archivo: %w", err)
	}

	// Inicializar el archivo si está vacío
	if len(dataUsers.Users) == 0 {
		if err := InitializeData(filename, &dataUsers, userName); err != nil {
			return fmt.Errorf("error al inicializar el archivo: %w", err)
		}
		return nil
	}

	// Verificar si el usuario ya existe
	for _, user := range dataUsers.Users {
		if user.Name == userName {

			fmt.Printf("El usuario %s ya existe.\n", user.Name)
			return nil
		}
	}

	// Si el usuario no existe, crear uno nuevo y lo agrega a la lista
	newUser := users.NewUser(userName)
	newUser.UserId = len(dataUsers.Users) + 1

	dataUsers.Users = append(dataUsers.Users, newUser)

	if err := WriteData(filename, dataUsers); err != nil {
		return fmt.Errorf("error al escribir el archivo: %w", err)
	}

	fmt.Println("Nuevo usuario agregado.")
	return nil
}

func FindUser(userName string, dataUsers InfoUsers) (users.User, error) {
	// Buscar el usuario en el archivo JSON
	for i := range dataUsers.Users {
		if dataUsers.Users[i].Name == userName {
			return dataUsers.Users[i], nil
		}
	}
	return users.User{}, fmt.Errorf("invalid user name")
}
