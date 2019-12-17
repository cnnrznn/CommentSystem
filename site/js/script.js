$(document).ready(function() {

$.get("http://localhost:8888/Comment/List?comment_id=0", function(data, status) {
    console.log(status);
    console.log(JSON.parse(data));
});

});
