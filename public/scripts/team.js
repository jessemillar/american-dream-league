function makeTeam() {
	for (var i in allTeams) { // Find the ID of the selected team
		if (allTeams[i].name == selectedTeam) {
			selectedTeam = allTeams[i].id;
		}
	}

	body = {
		team: selectedTeam.toString(),
		amount: $("#make-team-amount").val(),
		recipient: $("#make-team-recipient").val(),
		note: $("#make-team-note").val()
	};

	$.ajax("/api/team", {
		"data": JSON.stringify(body),
		"type": "POST",
		"processData": false,
		"contentType": "application/json",
		"success": function(data) {
			window.location.href = "/frontend#teams";
			location.reload();
		}
	});
}

function getPlayers(callback) {
	$.get("/api/player/team/1", function(data) {
		if (callback) {
			callback(data);
		} else {
			return data;
		}
	});
}

function populatePlayers(players) {
	for (var i in players) {
		var tr = document.createElement("tr");
		var name = document.createElement("td");
		var age = document.createElement("td");
		var position = document.createElement("td");

		name.appendChild(document.createTextNode(players[i].name.firstName + " " + players[i].name.middleName + " " + players[i].name.lastName));
		age.appendChild(document.createTextNode(players[i].age + "yrs"));
		position.appendChild(document.createTextNode(players[i].position.name));

		tr.appendChild(name);
		tr.appendChild(age);
		tr.appendChild(position);

		$("#team-page #players-list").append(tr);
	}
}
