package main

// import (
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// // type GameState struct {
// // 	StateName string
// // 	OnEnter   func()
// // 	OnUpdate  func()
// // 	OnExit    func()
// // }

// type MenuItem struct {
// 	Text       string
// 	PositionX  int32
// 	PositionY  int32
// 	IsSelected bool
// }

// func drawMenu(menuItems []MenuItem) {
// 	for _, item := range menuItems {
// 		if item.IsSelected {
// 			rl.DrawText(item.Text, item.PositionX, item.PositionY, 32, rl.DarkGray)
// 		} else {
// 			rl.DrawText(item.Text, item.PositionX, item.PositionY, 32, rl.LightGray)
// 		}
// 	}
// }

// func maine() {
// 	screenWidth := int32(800)
// 	screenHeight := int32(600)

// 	rl.InitWindow(screenWidth, screenHeight, "Menu Example")
// 	defer rl.CloseWindow()

// 	// Inicjalizuj odtwarzacz muzyki
// 	rl.InitAudioDevice()

// 	// Wczytaj muzykę
// 	music := rl.LoadMusicStream("path/to/your/music.ogg")
// 	rl.PlayMusicStream(music)

// 	menuItems := []MenuItem{
// 		{"Start", 100, 200, true},
// 		{"Options", 100, 250, false},
// 		{"Exit", 100, 300, false},
// 	}

// 	selectedItem := 0

// 	menuState := GameState{
// 		StateName: "MenuState",
// 		OnEnter: func() {
// 			// Inicjalizacja stanu menu
// 		},
// 		OnUpdate: func() {
// 			rl.ClearBackground(rl.RayWhite)

// 			if rl.IsKeyPressed(rl.KeyDown) {
// 				menuItems[selectedItem].IsSelected = false
// 				selectedItem = (selectedItem + 1) % len(menuItems)
// 				menuItems[selectedItem].IsSelected = true
// 			} else if rl.IsKeyPressed(rl.KeyUp) {
// 				menuItems[selectedItem].IsSelected = false
// 				selectedItem = (selectedItem - 1 + len(menuItems)) % len(menuItems)
// 				menuItems[selectedItem].IsSelected = true
// 			} else if rl.IsKeyPressed(rl.KeyEnter) {
// 				selectedOption := menuItems[selectedItem].Text
// 				switch selectedOption {
// 				case "Start":
// 					// Rozpocznij grę
// 					rl.Trace("Start game selected\n")
// 				case "Options":
// 					// Przejdź do stanu OptionsState
// 					rl.Trace("Options selected\n")
// 					currentGameState = &optionsState
// 					currentGameState.OnEnter()
// 				case "Exit":
// 					// Zakończ aplikację
// 					rl.Trace("Exit selected\n")
// 					rl.CloseWindow()
// 				}
// 			}

// 			drawMenu(menuItems)
// 		},
// 		OnExit: func() {
// 			// Czyszczenie stanu menu
// 		},
// 	}

// 	optionsState := GameState{
// 		StateName: "OptionsState",
// 		OnEnter: func() {
// 			// Inicjalizacja stanu opcji
// 		},
// 		OnUpdate: func() {
// 			rl.ClearBackground(rl.RayWhite)

// 			// Logika stanu opcji
// 			rl.DrawText("Options Screen", 100, 200, 32, rl.DarkGray)

// 			if rl.IsKeyPressed(rl.KeyEscape) {
// 				// Powróć do MenuState po naciśnięciu klawisza Escape
// 				currentGameState.OnExit()
// 				currentGameState = &menuState
// 				currentGameState.OnEnter()
// 			}
// 		},
// 		OnExit: func() {
// 			// Czyszczenie stanu opcji
// 		},
// 	}

// 	// Inicjalizuj początkowy stan gry jako MenuState
// 	currentGameState := &menuState
// 	currentGameState.OnEnter()

// 	for !rl.WindowShouldClose() {
// 		rl.BeginDrawing()

// 		currentGameState.OnUpdate()

// 		rl.EndDrawing()
// 	}

// 	// Zwolnij zasoby muzyki
// 	rl.UnloadMusicStream(music)

// 	// Zamknij odtwarzacz audio
// 	rl.CloseAudioDevice()
// }
