$(document).ready(function() {

$.get("http://localhost:8888/Comment/List?comment_id=0", function(data, status) {
    console.log(status, JSON.parse(data));

    json = JSON.parse(data);

    $('#comments').html(JSON.stringify(json, null, 4));
});

});
