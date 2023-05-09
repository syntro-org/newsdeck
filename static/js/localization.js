document.addEventListener("DOMContentLoaded", function() {  
  let matches = document.cookie.match(
    new RegExp(
      "(?:^|; )" +
        "lang".replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, "\\$1") +
        "=([^;]*)"
    )
  );

  let lang = decodeURIComponent(matches[1]);
  let fileName = document.getElementById("page-name").getAttribute("content")

  let xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      let data = JSON.parse(this.responseText);
      console.log(data);

      let elementsWithData = document.querySelectorAll(".json-data");
      for (const element of elementsWithData) {
        const fieldName = element.dataset.field;
        if (fieldName && data[fieldName]) {
          element.innerHTML = data[fieldName];
        }
      }
    }
  };
  xmlhttp.open(
    "GET",
    "static/content/" + lang + "/" + fileName + ".json",
    true
  );
  xmlhttp.send();
});