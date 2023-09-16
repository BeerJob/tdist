package main
import (
	"log"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"context"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	pb "github.com/BeerJob/tdist/proto"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	content, err := os.ReadFile("parametros_de_inicio.txt")
	if err != nil{
		log.Fatal("Fatal error")
	}
	rContent := strings.SplitN(string(content), "-",2)
	iLimit, err := strconv.Atoi(rContent[0])
	sLimit, err := strconv.Atoi(strings.SplitN(rContent[1], "\n", 2)[0])
	amount, err := strconv.Atoi(strings.SplitN(rContent[1], "\n", 2)[1])
	if amount==-1{
		amount = 1000000
	}
	for i:=1; i<=amount; i++{
		created := rand.Intn(sLimit-iLimit+1)+iLimit
		if amount==1000000{
			log.Printf("Generacion %d/infinito", i)
		}else{
			log.Printf("Generacion %d/%d", i, amount)
		}
		//Servidor 1
		conn, err := grpc.Dial("10.6.46.140:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con Servidor1!")
		}
		defer conn.Close()
		cliente := pb.NewServidorRegionalClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Print("No hay respuesta del Servidor1")
		}else{
			log.Printf("Respuesta sincrona del Servidor1: %s", r.Ok)
		}
		//Servidor2
		conn, err = grpc.Dial("10.6.46.141:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con Servidor2!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Print("No hay respuesta del Servidor2")
		}else{
			log.Printf("Respuesta sincrona del Servidor2: %s", r.Ok)
		}
		//Servidor3
		conn, err = grpc.Dial("10.6.46.109:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con Servidor3!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Print("No hay respuesta del Servidor3")
		}else{
			log.Printf("Respuesta sincrona del Servidor3: %s", r.Ok)
		}
		//Servidor4
		conn, err = grpc.Dial("10.6.46.110:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se pudo conectar con Servidor4!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Print("No hay respuesta del Servidor4")
		}else{
			log.Printf("Respuesta sincrona del Servidor4: %s", r.Ok)
		}
		recibido := 0
		noinscritos := 0
		timer := time.NewTicker(10*time.Second)
		/*
		//Servidor1
		connentionRabbitMQ, err := amqp.Dial("ampq//guest:guest@10.6.46.109:8082")
		if err != nil{
			log.Print("No se pudo conectar a la cola")
		}
		defer connentionRabbitMQ.Close()
		channel, err := connentionRabbitMQ.Channel()
		if err != nil{
			log.Print("No se pudo crear canal en la cola")
		}
		defer channel.Close()
		timer = time.NewTicker(10* time.Second)
		for{
			select{
			case <- timer.C:
				msg, ok, err := channel.Get("centralQueue", true)
				if err != nil{
					log.Print("Error al obtener mensaje de la cola")
				}
				if ok{
					log.Printf("Mensaje asincrono de Servidor1 leido")
					recibido, err = strconv.Atoi(string(msg.Body))
				}
			}
		}
		
		if created-recibido < 0{
			noinscritos = -(created-recibido)
			created=0
		}else{
			noinscritos=0
			created = created-recibido
		}
		log.Printf("Se inscribieron %d cupos del Servidor1", recibido-noinscritos)
		conn, err = grpc.Dial("10.6.46.140:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se puede conectar con Servidor1")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(noinscritos)})
		if err != nil{
			log.Print("No hay respuesta del Servidor1")
		}

		//Servidor2
		connentionRabbitMQ, err = amqp.Dial("ampq//guest:guest@10.6.46.109:8082")
		if err != nil{
			log.Print("No se pudo conectar a la cola")
		}
		defer connentionRabbitMQ.Close()
		channel, err = connentionRabbitMQ.Channel()
		if err != nil{
			log.Print("No se pudo crear canal en la cola")
		}
		defer channel.Close()
		timer = time.NewTicker(10* time.Second)
		for{
			select{
			case <- timer.C:
				msg, ok, err := channel.Get("centralQueue", true)
				if err != nil{
					log.Print("Error al obtener mensaje de la cola")
				}
				if ok{
					log.Printf("Mensaje asincrono de Servidor1 leido")
					recibido, err = strconv.Atoi(string(msg.Body))
				}
			}
		}
		if created-recibido < 0{
			noinscritos = -(created-recibido)
			created=0
		}else{
			noinscritos=0
			created = created-recibido
		}
		log.Printf("Se inscribieron %d cupos del Servidor2", recibido-noinscritos)
		conn, err = grpc.Dial("10.6.46.141:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se puede conectar con Servidor2")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(noinscritos)})
		if err != nil{
			log.Print("No hay respuesta del Servidor2")
		}
		*/
		//Servidor3
		connentionRabbitMQ, err := amqp.Dial("ampq//guest:guest@10.6.46.109:8082")
		if err != nil{
			log.Print("No se pudo conectar a la cola")
		}
		defer connentionRabbitMQ.Close()
		channel, err := connentionRabbitMQ.Channel()
		if err != nil{
			log.Print("No se pudo crear canal en la cola")
		}
		defer channel.Close()
		timer = time.NewTicker(10* time.Second)
		for{
			select{
			case <- timer.C:
				msg, ok, err := channel.Get("centralQueue", true)
				if err != nil{
					log.Print("Error al obtener mensaje de la cola")
				}
				if ok{
					log.Printf("Mensaje asincrono de Servidor1 leido")
					recibido, err = strconv.Atoi(string(msg.Body))
				}
			}
		}
		if created-recibido < 0{
			noinscritos = -(created-recibido)
			created=0
		}else{
			noinscritos=0
			created = created-recibido
		}
		log.Printf("Se inscribieron %d cupos del Servidor3", recibido-noinscritos)
		conn, err = grpc.Dial("10.6.46.109:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se puede conectar con Servidor3")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(noinscritos)})
		if err != nil{
			log.Print("No hay respuesta del Servidor3")
		}
		/*
		//Servidor4
		connentionRabbitMQ, err = amqp.Dial("ampq//guest:guest@10.6.46.109:8082")
		if err != nil{
			log.Print("No se pudo conectar a la cola")
		}
		defer connentionRabbitMQ.Close()
		channel, err = connentionRabbitMQ.Channel()
		if err != nil{
			log.Print("No se pudo crear canal en la cola")
		}
		defer channel.Close()
		timer = time.NewTicker(10* time.Second)
		for{
			select{
			case <- timer.C:
				msg, ok, err := channel.Get("centralQueue", true)
				if err != nil{
					log.Print("Error al obtener mensaje de la cola")
				}
				if ok{
					log.Printf("Mensaje asincrono de Servidor1 leido")
					recibido, err = strconv.Atoi(string(msg.Body))
				}
			}
		}
		if created-recibido < 0{
			noinscritos = -(created-recibido)
			created=0
		}else{
			noinscritos=0
			created = created-recibido
		}
		log.Printf("Se inscribieron %d cupos del Servidor4", recibido-noinscritos)
		conn, err = grpc.Dial("10.6.46.110:8080", grpc.WithInsecure())
		if err != nil{
			log.Print("No se puede conectar con Servidor4")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(noinscritos)})
		if err != nil{
			log.Print("No hay respuesta del Servidor4")
		}
		*/
	}
}