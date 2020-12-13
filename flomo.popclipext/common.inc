<?php

function err($str) {
	fwrite(STDERR, rtrim($str) . "\n");
}