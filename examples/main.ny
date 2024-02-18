gawo zamoyo {
    gulu Nyama {
        ndondomeko yenda() {
            // Yenda...
        }
        
        ndondomeko puma () {
            // Puma...
        }
    }

    gulu Munthu < Nyama {
        ndondomeko yenda () {
            kholo.yenda();	// Panga ndondomeko yoyenda ya nyama
        }
    }

    munthu = Munthu();
    ngati munthu ali Nyama {
        nena 'Munthu ndi nyama';
    } kapena {
        nena 'Munthu si nyama';
    }
}