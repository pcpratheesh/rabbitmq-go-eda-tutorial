const socket = new WebSocket('ws://localhost:8080/ws');

socket.addEventListener('open', function (event) {
  console.log('WebSocket connection opened');
});

socket.addEventListener('message', function (event) {
    const data = JSON.parse(event.data)
    switch (data.type) {
        case "consumer" : 
            AppendConsumer(data.name)
        break
        default:
                console.log("Data type : ",data.type)
        break;
    }
  console.log('Received message:', event.data);
});

socket.addEventListener('error', function (event) {
  console.error('WebSocket error:', event);
});

socket.addEventListener('close', function (event) {
  console.log('WebSocket connection closed');
});

function sendMessage(message) {
  socket.send(message);
}



// AppendConsumer
function AppendConsumer(name){
    $(".consumers-list").append(`
        <div class="d-block p-2">
            <div class="card">
            <div class="card-body">
                <div class="lead">`+name+`</div>
            </div>
            </div>
        </div>
    `)

}