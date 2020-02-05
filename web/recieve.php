<?php

include "common.php";

error_log($_SERVER['REQUEST_URI']);
$parts = parse_url($_SERVER['REQUEST_URI']);
parse_str($parts['query'], $query);
// echo $query['email'];

if (count($_GET) == 0) {
	echo "No data received.";
	exit;
}

// if (!in_array($_SERVER['REMOTE_ADDR'], $ipGameServer)) die();

if (!file_exists("./data/data.db")) {
    error_log("Database not found. Please ensure you have ran the installer first.");
	echo "Database not found. Please ensure you have ran the installer first.";
	exit;
}
// echo $query['email'];
$option = $query['option'];
//echo "Raw data received: " .implode(" ", $_GET) . "\n";
$string = "";
foreach($query as $key => $value) {
	$string .= $key . ": " . $value . "\n";
}

// Truncate string if too large
if (strlen($string) > 500) {
	$string = substr($string, 0, 500) . "...";
}
error_log ("Processed data: \n". $string . "\n");

function str2filename($str)
{
	// TODO: Change to the correct regular expression
	$res = preg_replace('/[\0\r\n\\\"\'\/\[\]:;\|=\?\<\>\*\.\,]/', '', $str);
	return $res;
}

if ($option == "addFile") { // Add receieved file to this directory
    $fileName = $query["fileName"];
	// $fileContents = $_FILES["fileContents"];
    
	error_log('FileName: '.$fileName);
    $content = file_get_contents('php://input');
	$fp = fopen("./data/".$fileName.".json", "w");
	fwrite($fp, $content);
	fclose($fp);
} elseif ($option == "dbInsert") { // Insert values into SQLite database
	$date = date("Y-m-d");
	$serverId = -1;
	try {
		$db = new PDO('sqlite:data/data.db');

		$sql = "INSERT INTO operations 
			(world_name, mission_name, mission_duration, filename, date, type) VALUES (
				:world_name,
				:mission_name,
				:mission_duration,
				:filename,
				:date,
				:type
			)
		";
		$sth = $db->prepare($sql, array(PDO::ATTR_CURSOR => PDO::CURSOR_FWDONLY));
		$sth->execute(array(
			':world_name' => $_GET["worldName"],
			':mission_name' => $_GET["missionName"],
			':mission_duration' => $_GET["missionDuration"],
			':filename' => str2filename($_GET["filename"]),
			':date' => $date,
			':type' => $_GET["type"]
		));
		
		// TODO: Increment local capture count

		// Get server ID
		$results = $db->query("SELECT remote_id FROM servers");
		$serverId = $results->fetch()["remote_id"];
		$db = null;
		print_debug($serverId);

	} catch (PDOExecption $e) {
		error_log("Exception: ".$e->getMessage());
	}

	// TODO: Increment capture count on remote database
	$result = curlRemote("stats-manager.php", array(
		"option" => "increment_capture_count",
		"server_id" => $serverId
	));

    print_debug($result);
} else {
	echo "Invalid option";
}

?>
