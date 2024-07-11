package services

import (
	"math/rand"
	"module/internal/models"
	"strconv"
)

///
///
/// ТУТ НАДО ПРОВЕРИТЬ ЧТО devNum не равно rands. вроде же нужно делать обновление перед каждым выводом? или только один раз?
/// rand.Seed(time.Now().UnixNano())
///

func EnrichtedOwner() []models.People {

	var nameMap = map[int]string{1: "Ivan", 2: "Sergey", 3: "Oleg", 4: "Anton", 5: "Dima", 6: "Sasha", 7: "Nikita", 8: "Slava", 9: "Kostya"}
	var surnMap = map[int]string{1: "Ivanov", 2: "Smirnov", 3: "Sidorov", 4: "Kochkin", 5: "Subarov", 6: "Nikiforov", 7: "Osipov", 8: "Vasin", 9: "Petrov"}
	var mailMap = map[int]string{1: "@yandex", 2: "@mail", 3: "@gmail", 4: "@yahoo", 5: "@bing", 6: "@outlook"}

	var peoples []models.People

	var devNum = randRange(1, 3)
	for i := 0; i < devNum; i++ {
		var curUser models.People
		curUser.Name = nameMap[randRange(1, len(nameMap))]
		curUser.Surname = surnMap[randRange(1, len(surnMap))]
		curUser.Email = curUser.Name + strconv.Itoa(randRange(1800, 2000)) + mailMap[randRange(1, len(mailMap))]

		peoples = append(peoples, curUser)
	}

	return peoples
}

// рандомим количество устройств, и суём радомные в массив
func EnrichtedDevices() []models.AdditionalDevices {

	var devMap = map[int]string{1: "ParkMaster", 2: "Trailer coupling", 3: "Window lifter", 4: "Power steering", 5: "Nitro", 6: "DVR"}

	var devNum = randRange(1, 3)
	var rands = randRange(1, 3)
	var curRand = randRange(1, 6)

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

// рандомим значения для двигателя
func EnrichtedEngine() models.CarEngine {

	var engine models.CarEngine
	engine.EngineCapacity = float64(randRange(500, 2000))
	engine.EnginePower = float64(randRange(100, 600))

	return engine
}

// функция рандома от и до
func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
