var yes = "success";
var no = "fail";
var counter = 0;
var answer1 = "button3";
var answer2 = "buttonX";
var answer3 = "button5";
var answer4 = "buttonX";

// The initializor

onEvent("initButton", "click", function( ) {
  setScreen("firstQuestion");
});

function success() { 
   setScreen(yes);
   counter += 1;
}

// Handling the first scene

function questionX(buttonx) {
   onEvent(buttonx, "click", function() {
      if (q1 === answer1) {
         success()
      }
      else {
         setScreen(no);
      }
   });
   return counter;
}

// The fail case

onEvent("button3", "click", function() {
});

