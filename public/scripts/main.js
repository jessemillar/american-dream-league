var currentUser = "";

function init() {
	if (window.location.hash) {
		var page = window.location.hash.substr(1);
		setCurrentPage(page);
	} else {
		setCurrentPage("login");
	}

    getTeams(populateTeams);
    getPlayers(populatePlayers);
}

function setCurrentPage(page) {
	bootpage.show(page);
	setActiveNavigation(page);
}

function setActiveNavigation(link) {
	var defaultClass = "four columns";

	// Reset all links
	document.getElementById("leagues-page-navigation").className = defaultClass;
	document.getElementById("profile-page-navigation").className = defaultClass;
	document.getElementById("settings-page-navigation").className = defaultClass;

	// Make the link we care about active
	document.getElementById(link + "-page-navigation").className += " active";
}
