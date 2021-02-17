package sayiTahmin

import (
	"fmt"
	"math/rand"
	"time"
)

func Game() {
	rand.Seed(time.Now().Unix())
	aklimdakiSayi := rand.Intn(100)
	tahminEdilen := 0

	tahminSayısı := 0
	for aklimdakiSayi != tahminEdilen {
		fmt.Print("Bir sayı giriniz: ")
		fmt.Scanln(&tahminEdilen)
		if tahminEdilen >= 1 && tahminEdilen <= 100 {
			tahminSayısı++
			if tahminEdilen > aklimdakiSayi {
				fmt.Println("Daha küçük bir sayı giriniz.")
			} else if tahminEdilen < aklimdakiSayi {
				fmt.Println("Daha büyük bir sayı giniz.")
			} else if tahminEdilen == aklimdakiSayi {
				if tahminSayısı <= 3 {
					fmt.Printf("Bravo Bildiniz. %v Tahmin : Süper\n", tahminSayısı)
				} else if tahminSayısı >= 4 && tahminSayısı <= 6 {
					fmt.Printf("Bravo Bildiniz. %v Tahmin : Güzel\n", tahminSayısı)
				} else if tahminSayısı >= 6 && tahminSayısı <= 8 {
					fmt.Printf("Bravo Bildiniz. %v Tahmin : Fena Değil\n", tahminSayısı)
				} else if tahminSayısı >= 8 && tahminSayısı <= 10 {
					fmt.Printf("Bravo Bildiniz. %v Tahmin : Kötü\n", tahminSayısı)
				} else if tahminSayısı > 10 {
					fmt.Printf("Bravo Bildiniz. %v Tahmin : Çok Kötü\n", tahminSayısı)
				}
			}
		} else if tahminEdilen <= 1 || tahminEdilen >= 100 {
			fmt.Println("Lütfen 1 ve 100 arasında bir sayı giriniz.")
		}

	}
}
