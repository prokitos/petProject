package services

import (
	"math/rand"
	"module/internal/models"
	"strconv"
)

// рандомная генерация машин
func EnrichtedBase() models.Car {

	var curCar models.Car

	var colorMap = map[int]string{1: "Red", 2: "White", 3: "Yellow", 4: "Green", 5: "Black", 6: "Blue"}
	var markMap = map[int]string{1: "Lada", 2: "Ford", 3: "BMW", 4: "Audi", 5: "Mazda", 6: "Toyota"}
	var seatMap = map[int]int{1: 2, 2: 4, 3: 6}

	curCar.Color = colorMap[randRange(1, len(colorMap)+1)]
	curCar.Mark = markMap[randRange(1, len(markMap)+1)]
	curCar.Year = strconv.Itoa(randRange(1970, 2020))
	curCar.Price = randRange(500000, 4000000)
	curCar.MaxSpeed = randRange(90, 260)
	curCar.SeatsNum = seatMap[randRange(1, len(seatMap)+1)]

	return curCar

}

// рандомная генерация владельцев
func EnrichtedOwner() []models.People {

	var nameMap = map[int]string{1: "Ivan", 2: "Sergey", 3: "Oleg", 4: "Anton", 5: "Dima", 6: "Sasha", 7: "Nikita", 8: "Slava", 9: "Kostya"}
	var surnMap = map[int]string{1: "Ivanov", 2: "Smirnov", 3: "Sidorov", 4: "Kochkin", 5: "Subarov", 6: "Nikiforov", 7: "Osipov", 8: "Vasin", 9: "Petrov"}
	var mailMap = map[int]string{1: "@yandex", 2: "@mail", 3: "@gmail", 4: "@yahoo", 5: "@bing", 6: "@outlook"}

	var peoples []models.People

	var devNum = randRange(1, 4) // количество владельцев от 1 до 3
	for i := 0; i < devNum; i++ {
		var curUser models.People
		curUser.Name = nameMap[randRange(1, len(nameMap)+1)]
		curUser.Surname = surnMap[randRange(1, len(surnMap)+1)]
		curUser.Email = curUser.Name + strconv.Itoa(randRange(1800, 2000)) + mailMap[randRange(1, len(mailMap)+1)]

		peoples = append(peoples, curUser)
	}

	return peoples
}

// рандомим количество устройств, и рандомное количество устрйств
func EnrichtedDevices() []models.AdditionalDevices {

	var devMap = map[int]string{1: "ParkMaster", 2: "Trailer coupling", 3: "Window lifter", 4: "Power steering", 5: "Nitro", 6: "DVR"}

	var devNum = randRange(1, 4)              // от 1 до 3 устройств
	var rands = randRange(1, 4)               // как будут генерироваться разные устройства
	var curRand = randRange(1, len(devMap)+1) // айди устройства в текущей генерации

	// супер странная генерация рандома
	var devices []models.AdditionalDevices
	for i := 0; i < devNum; i++ {

		var curDev models.AdditionalDevices
		curDev.DeviceName = devMap[curRand]
		devices = append(devices, curDev)

		if rands == 1 {
			curRand++

			if curRand > len(devMap) {
				curRand = 1
			}
		}
		if rands == 2 {
			curRand--
			if curRand < 1 {
				curRand = len(devMap)
			}
		}
		if rands == 3 {
			curRand += 2
			if curRand > len(devMap) {
				curRand = 2
			}
		}
	}

	return devices

}

// рандомная генерация двигателя
func EnrichtedEngine() models.CarEngine {

	var engine models.CarEngine
	engine.EngineCapacity = float64(randRange(500, 2000))
	engine.EnginePower = float64(randRange(100, 600))

	return engine
}

// функция рандома от и до  (до не включительно, поэтому нужно len + 1)
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
