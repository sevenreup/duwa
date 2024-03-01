nambala phatikiza(yambi: nambala, chiwiri: nambala) {
	bweza yamba + chiwiri;
}


ndondomeko lolowera() {
    lemba("Poyambira");

	nambala yoyamba = 6;
	nambala yachiwiri = 8;

	nambala yobwereza = phatikiza(yoyamba, yachiwiri);

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