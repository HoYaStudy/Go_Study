package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HoYaStudy/Go_Study/hcoin/blockchain"
	"github.com/HoYaStudy/Go_Study/hcoin/p2p"
	"github.com/HoYaStudy/Go_Study/hcoin/utils"
	"github.com/HoYaStudy/Go_Study/hcoin/wallet"
	"github.com/gorilla/mux"
)

type url string

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addTxPayload struct {
	To     string
	Amount int
}

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

type myWalletResponse struct {
	Address string `json:"address"`
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type addPeerPayload struct {
	Address, Port string
}

var port string

func Start(inPort int) {
	port = fmt.Sprintf(":%d", inPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status)
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	router.HandleFunc("/balance/{address}", balance)
	router.HandleFunc("/mempool", mempool)
	router.HandleFunc("/transactions", transactions).Methods("POST")
	router.HandleFunc("/wallet", myWallet).Methods("GET")
	router.HandleFunc("/ws", p2p.Upgrade).Methods("GET")
	router.HandleFunc("/peers", peers).Methods("GET", "POST")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

func (u urlDescription) String() string {
	return "Hello I'm the URL Description"
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application.json")
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the status of the Blockchain",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "GET: See All Blocks | POST: Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "Get TxOuts for an Address",
		},
		{
			URL:         url("/ws"),
			Method:      "GET",
			Description: "Upgrade to WebSocket",
		},
	}
	utils.HandleErr(json.NewEncoder(rw).Encode(data))
}

func status(rw http.ResponseWriter, r *http.Request) {
	blockchain.Status(blockchain.Blockchain(), rw)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blocks(blockchain.Blockchain())))
	case "POST":
		newBlock := blockchain.Blockchain().AddBlock()
		p2p.BroadcastNewBlock(newBlock)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.BalanceByAddress(address, blockchain.Blockchain())
		json.NewEncoder(rw).Encode(balanceResponse{address, amount})
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(address, blockchain.Blockchain())))
	}
}

func mempool(rw http.ResponseWriter, r *http.Request) {
	utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Mempool().Txs))
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	tx, err := blockchain.Mempool().AddTx(payload.To, payload.Amount)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(errorResponse{err.Error()})
		return
	}
	p2p.BroadcastNewTx(tx)
	rw.WriteHeader(http.StatusCreated)
}

func myWallet(rw http.ResponseWriter, r *http.Request) {
	address := wallet.Wallet().Address
	json.NewEncoder(rw).Encode(myWalletResponse{Address: address})
}

func peers(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(p2p.AllPeers(&p2p.Peers))
	case "POST":
		var payload addPeerPayload
		json.NewDecoder(r.Body).Decode(&payload)
		p2p.AddPeer(payload.Address, payload.Port, port, true)
		rw.WriteHeader(http.StatusOK)
	}
}
