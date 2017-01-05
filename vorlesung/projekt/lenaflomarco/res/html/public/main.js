function checkGetParameters() {
    location.search.substr(1).split("&").forEach(function(item) {
        tmp = item.split("=");
        if(tmp[0] === "status") {
            switch(tmp[1]) {
                case "failed":
                    alert("Login fehlgeschlagen.");
                    break;
                case "passwordsNotEqual":
                    alert("Die eingegebenen Kennwörter sind nicht identisch.");
                    break;
                case "userAlreadyExists":
                    alert("E-Mail Adresse bereits vergeben.");
                    break;
                case "error":
                    alert("Es ist ein Fehler aufgetreten");
                    break;
                case "badrequest":
                    alert("Es ist ein Fehler aufgetreten");
                    break;
                case "logout":
                    alert("Erfolgreich ausgeloggt.");
                    break;
                case "oldPasswordNotValid":
                    alert("Aktuelles Kennwort ungültig.");
                    break;
            }
        }
    })
}

checkGetParameters();