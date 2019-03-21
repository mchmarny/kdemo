$(function () {


    $(".logo img").click(function (event) {
        event.preventDefault();
        $("#logo-result-text").html("")
        var imgUrl = $(this).attr("src");
        $("#logo-url").val(imgUrl);
    });


    $("#logo-button").click(function (event) {
        event.preventDefault();
        var imgUrl = $("#logo-url").val();
        $.get("/logo?imageUrl=" + imgUrl, function (data) {
            console.log(data);
            $("#logo-result-text").html("<b>Result:</b> " + data.desc);
        });
    });


});