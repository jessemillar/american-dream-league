var bootpage = {
	currentPage: ""
};

bootpage.show = function(page, callback) {
	if (bootpage.currentPage) {
		document.getElementById(bootpage.currentPage + "-page").style.display = "none"; // Hide the current page
	}

	document.getElementById(page + "-page").style.display = "block"; // Show the desired page
	bootpage.currentPage = page;

	if (callback) {
		callback();
	}
};
