<?php
require 'common.inc';

$api = getenv('POPCLIP_OPTION_API');
$tag = getenv('POPCLIP_OPTION_TAG');
$content = getenv('POPCLIP_TEXT');
$browser_title = getenv("POPCLIP_BROWSER_TITLE");
$browser_url = getenv("POPCLIP_BROWSER_URL");

// execute request
$ch = curl_init($api);
$header = array("Content-Type: application/json");

if ($browser_title != '' && $browser_url != '') {
  $content .= "\n\n网页标题：{$browser_title}";
  $content .= "\nURL：{$browser_url}";
}

if ($tag != '') {
  $content .= "\n\n#{$tag}";
}

$postjson = json_encode(["content" => $content]);

curl_setopt($ch, CURLOPT_TIMEOUT, 10);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, TRUE);
curl_setopt($ch, CURLOPT_HTTPHEADER, $header);
curl_setopt($ch, CURLOPT_POST, 1);
curl_setopt($ch, CURLOPT_POSTFIELDS, $postjson);
$response = curl_exec($ch);

$code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
err("bitly shorten code: " . $code);

if ($code == 200 || $code == 201) { // 201 created
	$json = json_decode($response);
	$result = $json->message;
	echo $result;
	exit(0); // success
}  
else if ($code == 403) {
	exit(2); // bad auth
}
else {
	exit(1); // other error
}

?>