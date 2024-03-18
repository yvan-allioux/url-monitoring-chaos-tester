<?php

// Récupérer l'heure actuelle
$heure_actuelle = date('i');

// Vérifier si nous sommes dans les 30 premières minutes de l'heure
if ($heure_actuelle < 30) {
    // Renvoie une réponse JSON avec le type "success"
    $reponse = array(
        "Type" => "success",
        "Content" => "success we are in the 30 first minutes of the hour."
    );
} else {
    // Renvoie une réponse JSON avec le type "error"
    $reponse = array(
        "Type" => "error",
        "Content" => "error we are NOT in the 30 first minutes of the hour."
    );
}

// Envoie la réponse JSON au client
header('Content-Type: application/json');
echo json_encode($reponse);

?>
