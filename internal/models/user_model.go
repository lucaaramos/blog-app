//put package

package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	UserName  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	Role      string    `json:"role" bson:"role"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Status    bool      `json:"status" bson:"status"`
}

func NewUser(name, email, username, password, role string) *User {
	return &User{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     email,
		UserName:  username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      role,
		Status:    true,
	}
}

// esta funcion se encarga de convertir la contraseña del usuario en un hash
// seguro antes de almacenarla en la base de datos.
func (u *User) HashPassword() error { //el receptor es un puntero a la estructura User, lo que permite modificar la estructura directamente
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost) //convierte la cadena en un slice de bytes y luego genera un hash
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword) //reemplaza la constraseña original del usuariopor el hash generado.
	// ahora la propiedad Password contieneel hash en lugar de la constraseña en texto plano
	return nil
}

// esta funcion se encarga de comparar la constraña proporcionada por el user con la hasheada almacenada.
func (u *User) CheckPassword(password string) error { //esta funcion recibe como argumento a password que es la constraseña que establecio el user durante el inicio de sesion o autenticacion
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err
	// if err == bcrypt.ErrMismatchedHashAndPassword {
	//     return errors.New("invalid password")
	// } else if err!= nil {
	//     return err
	// }
	// return nil
	// }
	// return errors.New("user not found")
	//
}
