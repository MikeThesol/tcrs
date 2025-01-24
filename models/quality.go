package models

import (
	rand2 "math/rand"
)

type Quality interface {
	GetMessageOfQuality() string
}

type GoodQuality struct {
	Message string
}

type BadQuality struct {
	Message string
}

func (g GoodQuality) GetMessageOfQuality() string {
	return g.Message
}

func (b BadQuality) GetMessageOfQuality() string {
	return b.Message
}

// Good quality
const (
	TheFoolGood           = "Начало чего-то нового, прыжок веры, потенциал, свободный дух."
	TheMagicianGood       = "Мастерство, способность вызывать нужное, творчество, сила воли."
	TheHighPriestessGood  = "Интуиция, бессознательное знание, руководство свыше."
	TheEmpressGood        = "Плодородие, изобилие, успех в делах, женственность."
	TheEmperorGood        = "Стабильные отношения, основанные на взаимном уважении."
	TheHierophantGood     = "Доверие и улучшение моральных ценностей."
	TheLoversGood         = "Любовь, гармония и выбор между двумя путями."
	TheChariotGood        = "Контроль над ситуацией, победа через усилия."
	TheStrengthGood       = "Сила духа и терпение, способность преодолевать трудности."
	TheHermitGood         = "Поиск истины и внутреннего света."
	TheWheelOfFortuneGood = "Циклы жизни и удача."
	TheJusticeGood        = "Справедливость и честность."
	TheHangedManGood      = "Жертва ради высшей цели, новый взгляд на вещи."
	TheDeathGood          = "Трансформация и новые начала."
	TheTemperanceGood     = "Гармония и баланс в жизни."
	TheDevilGood          = "Привязанность к материальному миру, искушения."
	TheTowerGood          = "Внезапные изменения и разрушение старых структур."
	TheStarGood           = "Надежда, вдохновение и исцеление."
	TheMoonGood           = "Иллюзии, интуиция и подсознание."
	TheSunGood            = "Счастье, успех и радость."
	TheJudgmentGood       = "Возрождение и принятие решений."
	TheWorldGood          = "Завершение цикла и достижение целей."
)

const (
	TheFoolBad           = "Страх двигаться вперед, наивность, безрассудство."
	TheMagicianBad       = "Манипуляции, обман, нарциссизм."
	TheHighPriestessBad  = "Заблокированная интуиция, скрытые планы."
	TheEmpressBad        = "Застой в делах, деспотизм."
	TheEmperorBad        = "Авторитаризм, конфликты."
	TheHierophantBad     = "Сопротивление переменам."
	TheLoversBad         = "Разногласия и трудности в отношениях."
	TheChariotBad        = "Потеря контроля и неуверенность."
	TheStrengthBad       = "Слабость и отсутствие уверенности." // Перевернутое: незавершенные дела и отсутствие удовлетворения.
	TheHermitBad         = "Изоляция и одиночество."
	TheWheelOfFortuneBad = "Несчастья и неудачи."
	TheJusticeBad        = "Предвзятость и несправедливость."
	TheHangedManBad      = "Бездействие и отсутствие прогресса."
	TheDeathBad          = "Страх перемен и стагнация."
	TheTemperanceBad     = "Дисбаланс и крайности."
	TheDevilBad          = "Освобождение от зависимостей."
	TheTowerBad          = "Избегание изменений и страх перед переменами."
	TheStarBad           = "Утрата надежды и недоверие к будущему."
	TheMoonBad           = "Заблуждения и страхи."
	TheSunBad            = "Недостаток уверенности и неудачи."
	TheJudgmentBad       = "Отказ от изменений и самокритика."
	TheWorldBad          = "Незавершенные дела и отсутствие удовлетворения."
)

func MakeQualityForCard(card *TarotCard) {
	rand := rand2.Intn(2)

	if rand == 0 {
		card.Quality = makeGoodQuality(card.Name)
		card.IsGood = true
	} else {
		card.Quality = makeBadQuality(card.Name)
		card.IsGood = false
	}

}

func makeGoodQuality(name string) GoodQuality {
	switch name {
	case "The Fool":
		return GoodQuality{TheFoolGood}
	case "The Magician":
		return GoodQuality{TheMagicianGood}
	case "The High Priestess":
		return GoodQuality{TheHighPriestessGood}
	case "The Empress":
		return GoodQuality{TheEmpressGood}
	case "The Emperor":
		return GoodQuality{TheEmperorGood}
	case "The Hierophant":
		return GoodQuality{TheHierophantGood}
	case "The Lovers":
		return GoodQuality{TheLoversGood}
	case "The Chariot":
		return GoodQuality{TheChariotGood}
	case "Strength":
		return GoodQuality{TheStrengthGood}
	case "The Hermit":
		return GoodQuality{TheHermitGood}
	case "Wheel of Fortune":
		return GoodQuality{TheWheelOfFortuneGood}
	case "Justice":
		return GoodQuality{TheJusticeGood}
	case "The Hanged Man":
		return GoodQuality{TheHangedManGood}
	case "Death":
		return GoodQuality{TheDeathGood}
	case "Temperance":
		return GoodQuality{TheTemperanceGood}
	case "The Devil":
		return GoodQuality{TheDevilGood}
	case "The Tower":
		return GoodQuality{TheTowerGood}
	case "The Star":
		return GoodQuality{TheStarGood}
	case "The Moon":
		return GoodQuality{TheMoonGood}
	case "The Sun":
		return GoodQuality{TheSunGood}
	case "Judgement":
		return GoodQuality{TheJudgmentGood}
	case "The World":
		return GoodQuality{TheWorldGood}
	default:
		return GoodQuality{}
	}
}

func makeBadQuality(name string) BadQuality {
	switch name {
	case "The Fool":
		return BadQuality{TheFoolBad}
	case "The Magician":
		return BadQuality{TheMagicianBad}
	case "The High Priestess":
		return BadQuality{TheHighPriestessBad}
	case "The Empress":
		return BadQuality{TheEmpressBad}
	case "The Emperor":
		return BadQuality{TheEmperorBad}
	case "The Hierophant":
		return BadQuality{TheHierophantBad}
	case "The Lovers":
		return BadQuality{TheLoversBad}
	case "The Chariot":
		return BadQuality{TheChariotBad}
	case "Strength":
		return BadQuality{TheStrengthBad}
	case "The Hermit":
		return BadQuality{TheHermitBad}
	case "Wheel of Fortune":
		return BadQuality{TheWheelOfFortuneBad}
	case "Justice":
		return BadQuality{TheJusticeBad}
	case "The Hanged Man":
		return BadQuality{TheHangedManBad}
	case "Death":
		return BadQuality{TheDeathBad}
	case "Temperance":
		return BadQuality{TheTemperanceBad}
	case "The Devil":
		return BadQuality{TheDevilBad}
	case "The Tower":
		return BadQuality{TheTowerBad}
	case "The Star":
		return BadQuality{TheStarBad}
	case "The Moon":
		return BadQuality{TheMoonBad}
	case "The Sun":
		return BadQuality{TheSunBad}
	case "Judgement":
		return BadQuality{TheJudgmentBad}
	case "The World":
		return BadQuality{TheWorldBad}
	default:
		return BadQuality{}
	}
}
