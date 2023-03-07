document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('count-button').addEventListener('click', function() {
      chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
        var tab = tabs[0];
        var url = new URL(tab.url);
        var path = url.pathname.split('/');
        var owner = path[1];
        var repo = path[2];
        var api_url = 'https://api.github.com/repos/' + owner + '/' + repo + '/contents';
        
        fetch(api_url)
          .then(response => response.json())
          .then(data => {
            var lines_of_code = {};
            
            for (var i = 0; i < data.length; i++) {
              var file = data[i];
              if (file.type == 'file') {
                var filename = file.name;
                var extension = filename.split('.').pop();
                if (lines_of_code[extension] == undefined) {
                  lines_of_code[extension] = 0;
                }
                fetch(file.download_url)
                  .then(response => response.text())
                  .then(data => {
                    var count = data.split('\n').length;
                    lines_of_code[extension] += count;
                    updateResult(lines_of_code);
                  });
              }
            }
          });
      });
    });
  });
  
  function updateResult(lines_of_code) {
    var result = document.getElementById('result');
    result.innerHTML = '';
    for (var language in lines_of_code) {
      var count = lines_of_code[language];
      result.innerHTML += '<p>' + language + ': ' + count + ' lines of code</p>';
    }
  }
  