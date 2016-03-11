<?php

$input = file_get_contents('php://input');
if (strpos($input, 'fail')) {
	http_response_code(500);

	printf('I failed for you: %s', file_get_contents('php://input'));
} else {
	http_response_code(200);
	printf('o hello, you sent me: %s', file_get_contents('php://input'));
}

