document.addEventListener("DOMContentLoaded", function() {  
  let matches = document.cookie.match(
    new RegExp(
      "(?:^|; )" +
        "lang".replace(/([\.$?*|{}\(\)\[\]\\\/\+^])/g, "\\$1") +
        "=([^;]*)"
    )
  );

  let lang = decodeURIComponent(matches[1]);
  document.documentElement.lang = lang;
  let fileName = document.getElementById("page-name").getAttribute("content")

  let xmlhttp = new XMLHttpRequest();
  xmlhttp.onreadystatechange = function () {
    if (this.readyState == 4 && this.status == 200) {
      let data = JSON.parse(this.responseText);
      let elementsWithData = document.querySelectorAll(".json-data");
      for (const element of elementsWithData) {
        const fieldName = element.dataset.field;
        if (fieldName && data[fieldName]) {
          element.innerHTML = data[fieldName];
        }
      }

      document.title = "Newsdeck - " + data.header;
    }
  };
  xmlhttp.open(
    "GET",
    "static/content/" + lang + "/" + fileName + ".json",
    true
  );
  xmlhttp.send();

  let dropdown = document.getElementById("lang-dropdown");
  for (let index = 0; index < dropdown.options.length; index++) {
    const element = dropdown.options[index];
    if(element.value == document.documentElement.lang) {
      element.selected = true;
      break;
    }
  }

  dropdown.addEventListener("change", onLanguageChanged);
});

function onLanguageChanged() {
  let dropdown = document.getElementById("lang-dropdown");
  document.cookie = "lang=" + dropdown.options[dropdown.selectedIndex].value;
  window.location.reload();
}