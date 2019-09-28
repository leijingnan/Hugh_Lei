package select_client

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mongodb/Command_line"
	"mongodb/web_portal"
	"net/http"
	"bufio"
	"os"
	"mongodb/client"
	"mongodb/Read_csv"
)
func Select_client(){
	for true {
		fmt.Println("please input the client(web portal/client/command line/read csv)(退出为exit)")
		inputReader := bufio.NewReader(os.Stdin)
		COMMAND, err := inputReader.ReadString('\n')							 //输入客户端想要的形式：portal web/client/command line
		if err == nil {
			if COMMAND == "exit\n" {
				break
			}

			//web portal 的形式
			if COMMAND == "web portal\n" {
				router := mux.NewRouter()
				router.HandleFunc("/Assets", web_portal.GetAssets).Methods("GET") //如果是get型method且路径为/Assets，下面几个命令同理
				router.HandleFunc("/Assets/{id}", web_portal.GetAsset).Methods("GET")
				// Post handle function
				router.HandleFunc("/Assets/{id}", web_portal.PostAsset).Methods("POST")
				// Delete handle function:
				router.HandleFunc("/Assets/{id}", web_portal.DeleteAsset).Methods("DELETE")
				// 启动 API端口0618
				log.Fatal(http.ListenAndServe(":0618", router))
			}

			//client的形式
			if COMMAND == "client\n" {
				client.Client()
			}

			//command line的形式
			if COMMAND == "command line\n" {
				command_line.CheckExe2()
			}
			if COMMAND == "read csv\n" {
				csv_reading.CSV_READ()
			}
		}
	}
}
