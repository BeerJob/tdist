package main
import (
	"log"
	//"fmt"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"context"

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
		amount = 100000
	}
	for i:=1; i<=amount; i++{
		created := rand.Intn(sLimit-iLimit+1)+iLimit
		if amount==100000{
			log.Printf("Generacion %d/infinito", i)
		}else{
			log.Printf("Generacion %d/%d", i, amount)
		}
		conn, err := grpc.Dial("adr: '173.20.0.1:50051'", grpc.WithInsecure())
		if err != nil{
			panic("cannot connect with the server!")
		}
		defer conn.Close()
		cliente := pb.NewServidorRegionalClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Fatal("Todo mal")
		}
		log.Printf("Respuesta de mensaje sincrono: %s", r.Ok)
		/*
		conn, err = grpc.Dial("adr: 'ip2:50501'", grpc.WithInsecure())
		if err != nil{
			panic("cannot connect with the server!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposDisponibles(ctx, &pb.Cupo{Cupos: strconv.Itoa(created)})
		if err != nil{
			log.Fatal("Todo mal")
		}
		log.Printf("Respuesta de mensaje sincrono: %s", r.Ok)
		*/
		//Codigo de la cola rabbit
		recibido := 3
		inscritos := 0
		if amount-recibido < 0{
			inscritos = -(amount-recibido)
			amount = 0
		}else{
			amount = amount-recibido
		}
		conn, err = grpc.Dial("adr: '173.20.0.1:50051'", grpc.WithInsecure())
		if err != nil{
			panic("cannot connect with the server!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(inscritos)})
		if err != nil{
			log.Fatal("Todo mal")
		}
		log.Printf("Respuesta de mensaje sincrono: %s", r.Ok)
		//Codigo cola rabbit
		/*
		recibido = 3
		inscritos = 0
		if amount-recibido < 0{
			inscritos = -(amount-recibido)
			amount = 0
		}else{
			amount = amount-recibido
		}
		conn, err = grpc.Dial("adr: 'ip2:50501'", grpc.WithInsecure())
		if err != nil{
			panic("cannot connect with the server!")
		}
		defer conn.Close()
		cliente = pb.NewServidorRegionalClient(conn)
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err = cliente.CuposRechazados(ctx, &pb.Rechazado{Rechazados: strconv.Itoa(inscritos)})
		if err != nil{
			log.Fatal("Todo mal")
		}
		log.Printf("Respuesta de mensaje sincrono: %s", r.Ok)
		*/
	}
}