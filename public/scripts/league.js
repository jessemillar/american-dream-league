function getTeams(callback) {
	$.get("/api/league", function(data) {
		if (callback) {
			callback(data);
		} else {
			return data;
		}
	});
}

function populateTeams(teams) {
	for (var i in teams) {
		var tr = document.createElement("tr");
		var name = document.createElement("td");
		var teamCount = document.createElement("td");

		name.appendChild(document.createTextNode(teams[i].name));
		teamCount.appendChild(document.createTextNode(teams[i].teamCount));

		tr.appendChild(name);
		tr.appendChild(teamCount);

		$("#leagues-page #teams-list").append(tr);
	}
}
