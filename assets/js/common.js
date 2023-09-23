$(function() {
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')
    $(".dataflow-b").append('<div class="data"></div>')


    // consumers list append
    $(".consumers-list").append(prepareConsumersList())

    var consumerCount = 1
    $("#addNewConsumer").on("click", function(){
        $("#consumed-data-table").html("") // clear the consumer data list

        $.ajax({
            type: "POST",
            url: "/consumer",
            data: JSON.stringify({
                name : "Consumer "+consumerCount,
            }),
            success: function (response) {
                consumerCount++
            }
        });
    })


    $("#run-producer").on("click", function(){
        // clear the table
        $("#consumed-data-table").html("")
        $.ajax({
            type: "POST",
            url: "/producer",
            success: function (response) {
                
            }
        });
    })
})


// prepareConsumersList
const prepareConsumersList = () => {
    let html_tag = ""
   
    return html_tag;
}