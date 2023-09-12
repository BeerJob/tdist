package main
import (
	"log"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math/rand"
	"time"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/BeerJob/tdist/proto"

)
func main(){
	conn, err := grpc.Dial(target:"ip", grpc.WithInsecure())
	if err != nil{
		panic("cannot connect with the server!")
	}
	defer conn.Close()
	cliente := pb.NewServidorRegionalClient(conn)
	rand.Seed(time.Now().UnixNano())
	content, err := os.ReadFile("parametros_de_inicio.txt")
	if err != nil{
		log.Fatal("Fatal error")
	}
	rContent := strings.SplitN(string(content), "-",2)
	iLimit, err := strconv.Atoi(rContent[0])
	sLimit, err := strconv.Atoi(strings.SplitN(rContent[1], "\n", 2)[0])
	amount, err := strconv.Atoi(strings.SplitN(rContent[1], "\n", 2)[1])
	for i:=0; i<amount; i++{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		fmt.Println(rand.Intn(sLimit-iLimit+1)+iLimit)
	}
}