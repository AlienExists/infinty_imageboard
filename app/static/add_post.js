function submit() {
   var data = document.getElementById("POST_SEND").value;

   var x = new XMLHttpRequest();
   x.open("POST", "/api/v1/posts?post=" + data, true);
   x.onload = function () {
      alert(x.responseText);
   }
   x.send(null);
   var Myelement = document.querySelector('textarea[name="POST_SEND"]');
   Myelement.value = "";
   location.reload();
}
