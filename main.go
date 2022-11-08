// create a weather app

function weather(city) {
    var url = 'http://api.openweathermap.org/data/2.5/weather?q=' + city + '&units=imperial&APPID=e9b9f8f8d8f9f9f9f9f9f9f9f9f9f9f9';
    var request = new XMLHttpRequest();
    request.open('GET', url, true);
    request.onload = function() {
        var data = JSON.parse(this.response);
        if (request.status >= 200 && request.status < 400) {
            console.log(data);
            var city = data.name;
            var temp = data.main.temp;
            var humidity = data.main.humidity;
            var wind = data.wind.speed;
            var weather = data.weather[0].description;
            var icon = data.weather[0].icon;
            var iconurl = 'http://openweathermap.org/img/w/' + icon + '.png';
            var weather_div = document.getElementById('weather');
            weather_div.innerHTML = '<h2>' + city + '</h2>' + '<h2>' + temp + '&deg;F</h2>' + '<h2>' + humidity + '%</h2>' + '<h2>' + wind + 'mph</h2>' + '<img src="' + iconurl + '">' + '<h2>' + weather + '</h2>';
        } else {
            console.log('error');
        }
    }
    request.send();
}

