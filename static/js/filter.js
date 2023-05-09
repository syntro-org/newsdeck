function addEventListenerToHeader() {
    const button = document.getElementById('collapsible-header');
    const content = document.getElementById('collapsible-content');
    
    button.addEventListener('click', function() {
      content.classList.toggle('hidden');
    });
}

addEventListenerToHeader();