nambala phatikiza(yambi: nambala, chiwiri: nambala) {
	bweza yamba + chiwiri;
}


ndondomeko lolowera() {
    lemba("Poyambira");

	namba yoyamba = 6;
	namba yachiwiri = 8;

	lemba(yoyamba  + yachiwiri);
	lemba(yoyamba  - yachiwiri);
	lemba(yoyamba  * yachiwiri);
	lemba(yoyamba  / yachiwiri);
	nambala yobwereza = phatikiza(yoyamba, yachiwiri);
	mawu dzina = "wona";
	lemba(dzina + " " + yoyamba);
	lemba(yobwereza);

    ngati(yoyamba > yachiwiri) {
        lemba("yoyamba ndiyayikulu");
    } kapena {
        lemba("yachiwiri ndiyayikulu");
    }

    pamene(x>4) {
    // Do something
    x++;
    }
}