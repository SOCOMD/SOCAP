<?php

// ini_set('display_errors', 1);
// ini_set('display_startup_errors', 1);
error_reporting(E_ALL);
ini_set('log_errors', TRUE);
ini_set('error_log', '/var/tmp/php-error.log');
ini_set('date.timezone','Australia/Sydney');


$debug = false;
$appTitle = "SOCAP";
$appDesc = "Socomd Operation Capture And Playback";
$appAuthor = "";
$ipGameServer = ["195.88.209.214", "193.19.118.241","localhost"];
$allTypeGames = [["Все", ""], ["TvT", "tvt"], ["IF", "if"]];
$lang = "en";
const VERSION = "2.2.1";

// Please do not modify this as these stats help me get a job. Thank-you! :)
// $statServerUrl = "http://138.201.116.116/ocap/remote/";

// Send cURL request to remote server
// $url should not include '/ocap/remote/' and should not have a leading '/'
function curlRemote($url, $postFields = array()) {
/*
	global $statServerUrl;
	$url = $statServerUrl . $url;

	$curl = curl_init();

	// Include post data, if supplied
	if (count($postFields) > 0) {
		curl_setopt($curl, CURLOPT_POST, true);
		curl_setopt($curl, CURLOPT_POSTFIELDS, http_build_query($postFields));
	}

	curl_setopt($curl, CURLOPT_URL, $url);
	curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
	$result = curl_exec($curl);
	curl_close($curl);

	return $result;
*/
}

function print_debug($var) {
	global $debug;
	if ($debug) {
		echo "<b>DEBUG: </b>";
		print_r($var);
		echo "<br/>";
	}
}
?>
