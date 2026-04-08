package main

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"example.com/taskservice/internal/config"
	"example.com/taskservice/internal/domain/task"
	taskdomain "example.com/taskservice/internal/domain/task"
	postgresinfra "example.com/taskservice/internal/infrastructure/postgres"
	postgrestask "example.com/taskservice/internal/repository/postgres"
)

var titles = []string{
	"Провести утренний обход пациентов", "Заполнить истории болезни", "Назначить лабораторные анализы", "Проверить жизненные показатели", "Выписать рецепты на лекарства", "Провести вакцинацию детей", "Обновить карты диспансерного наблюдения", "Подготовить пациентов к операции", "Провести ЭКГ исследование", "Выдать больничные листы", "Проверить срок годности медикаментов", "Провести перевязку послеоперационных ран", "Собрать анамнез у новых пациентов", "Направить на инструментальную диагностику", "Провести санитарную обработку палат", "Проверить результаты МРТ", "Провести реанимационные мероприятия", "Оформить эпикризы по выписке", "Провести забор крови на анализ", "Заполнить журнал учёта процедур",
}

var descriptions = []string{
	"Обойти 15 пациентов в терапевтическом отделении, зафиксировать жалобы", "Внести данные о лечении в электронные истории болезни по стандарту", "Выдать направления на ОАК, биохимию и коагулограмму для 10 пациентов", "Измерить АД, пульс, сатурацию и температуру у всех поступивших", "Выписать антибиотики и противовирусные согласно назначениям врача", "Поставить вакцину от гриппа и гепатита B в прививочном кабинете", "Обновить статусы хронических больных в регистре диспансеризации", "Провести предоперационную подготовку, забор анализов и ЭКГ", "Снять и расшифровать ЭКГ у пациентов кардиологического отделения", "Оформить и выдать листки нетрудоспособности сотрудникам поликлиники", "Провести ревизию аптечки: просроченные препараты утилизировать", "Сделать перевязку 8 пациентам после полостных операций", "Собрать жалобы, аллергоанамнез и сопутствующие заболевания", "Направить на УЗИ, КТ и рентген по клиническим показаниям", "Обработать палаты бактерицидными лампами и дезрастворами", "Расшифровать и подшить в карты результаты магнитно-резонансной томографии", "Провести СЛР при остановке дыхания и кровообращения по алгоритму", "Написать выписные эпикризы для 5 пациентов с инфарктом", "Взять венозную кровь у 12 пациентов на гормоны и онкомаркеры", "Записать все выполненные инъекции и капельницы в процедурный журнал",
}

var statuses = []task.Status{
	"new", "in_progress", "done", "new", "in_progress", "done", "canceled", "new", "in_progress", "done", "new", "in_progress", "done", "canceled", "new", "in_progress", "done", "new", "in_progress", "done",
}

func main() {
	dsn := config.LoadConfig().DatabaseDSN
	pool, err := postgresinfra.Open(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	task := postgrestask.New(pool)

	err = generateTasks(20, task)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Данные заполнены")
}

func generateTasks(count int, task *postgrestask.Repository) error {
	for i := 0; i < count-1; i++ {
		taskData := &taskdomain.Task{
			Title:       titles[rand.IntN(len(titles))],
			Description: descriptions[rand.IntN(len(descriptions))],
			Status:      statuses[rand.IntN(len(statuses))],
			CreatedAt:   time.Now().Add(time.Hour * 24 * -time.Duration(rand.IntN(10))),
			UpdatedAt:   time.Now().Add(time.Hour * 24 * -time.Duration(rand.IntN(10))),
		}
		_, err := task.Create(context.Background(), taskData)
		if err != nil {
			return err
		}
	}
	return nil
}
