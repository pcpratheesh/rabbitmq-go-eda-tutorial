const eventSource = new EventSource('event');

eventSource.addEventListener('event', function (event) {
    const data = JSON.parse(event.data)
    
    switch (data.type) {
        case "consumer" : 
            AppendConsumer(data.name)
        break
        case "consumed-data" : 
            AppendConsumedData(data)
        break
        default:
            console.log("Data type : ",data.type)
        break;
    }

    if (data.completed == true) {
        closeEventStream()
    }
});

eventSource.onmessage = (event) => {
    const data = JSON.parse(event.data)
    console.log("on message : ",data)
}

eventSource.onerror = (error) => {
    console.error('SSE error:', error);
};

// Either this or the up one
eventSource.addEventListener('error', (error) => {
    console.error('SSE error:', error);
});

eventSource.onopen = () => {
    console.log('SSE connection opened');
};

eventSource.onclose = () => {
    console.log('SSE connection closed');
};

function closeEventStream(){
    console.log("event stream closed")
    eventSource.close() 
}


// AppendConsumer
function AppendConsumer(name){
    $(".consumers-list").append(`
        <div class="entry p-2">
            <div class="card">
            <div class="card-body">
                <div class="lead">`+name+`</div>
            </div>
            </div>
        </div>
    `)

}


function AppendConsumedData(data){
    $("#consumed-data-table").append(`
        <tr>
            <td><button class="btn">`+data.name+`</button></td>
            <td>`+data.data+`</td>
        </tr>
    `)
}

