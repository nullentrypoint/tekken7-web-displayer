// Code from: https://github.com/ParadiseAigo/Tekken-7-Player-Displayer/blob/master/Tekken-7-Player-Displayer/pointers.h

package tekken


const TEKKEN_WINDOW_NAME = "TEKKEN 7 "
const TEKKEN_EXE_NAME = "TekkenGame-Win64-Shipping.exe"
const TEKKEN_STEAM_APP_ID = "389730"

var ALL_CHARACTERS = [...]string{"Paul", "Law", "King", "Yoshimitsu", "Hwoarang", "Xiaoyu", "Jin", "Bryan", "Heihachi", "Kazuya", "Steve", "Jack-7", "Asuka", "Devil Jin", "Feng", "Lili", "Dragunov", "Leo", "Lars", "Alisa", "Claudio", "Katarina", "Lucky Chloe", "Shaheen", "Josie", "Gigas", "Kazumi", "Devil Kazumi", "Nina", "Master Raven", "Lee", "Bob", "Akuma", "Kuma", "Panda", "Eddy", "Eliza", "Miguel", "Soldier", "Kid Kazuya", "Jack-#", "Young Heihachi", "Dummy A", "Geese", "Noctis", "Anna", "Lei", "Marduk", "Armor King", "Julia", "Negan", "Zafina", "Ganryu", "Leroy Smith", "Fahkumram", "Kunimitsu", "Lidia"}

const STEAM_API_MODULE_NAME = "steam_api64.dll"
const STEAM_API_MODULE_EDITED_NAME = "steam_api64_o.dll"
const TEKKEN_MODULE_NAME = "TekkenGame-Win64-Shipping.exe"

// steam id of the opponent
const STEAM_ID_BETTER_STATIC_POINTER = 0x2FC50  //to get the real static pointer: needs to be added to the module name, example: "steam_api64_o.dll"+2FC50
var STEAM_ID_BETTER_POINTER_OFFSETS = [...]int{8, 0x10, 0, 0x10, 0x28, 0x88, 0}

// steam id of the user
const STEAM_ID_USER_STATIC_POINTER = 0x2FF78  //to get the real static pointer: needs to be added to the module name, example: "steam_api64_o.dll"+2FC50
var STEAM_ID_USER_POINTER_OFFSETS = [...]int{}

// steam name of opponent after having accepted
const OPPONENT_STRUCT_NAME_STATIC_POINTER = 0x34D2520  //to get the real static pointer: needs to be added to the module name
var OPPONENT_STRUCT_NAME_POINTER_OFFSETS = [...]int{0, 0x8, 0x11C}

// character of opponent after having accepted
const OPPONENT_STRUCT_CHARACTER_STATIC_POINTER = 0x34D2520  //to get the real static pointer: needs to be added to the module name
var OPPONENT_STRUCT_CHARACTER_POINTER_OFFSETS = [...]int{0, 0x8, 0x10}

const SECONDS_REMAINING_MESSAGE_STATIC_POINTER = 0x34CB960  //to get the real static pointer: needs to be added to the module name
var SECONDS_REMAINING_MESSAGE_POINTER_OFFSETS = [...]int{0x58, 0x330, 0x40, 0x30, 0x158, 0x750}

const SCREEN_MODE_STATIC_POINTER = 0x347ADF8  //to get the real static pointer: needs to be added to the module name
var SCREEN_MODE_POINTER_OFFSETS = [...]int{}

const SCREEN_MODE_FULLSCREEN = 0
const SCREEN_MODE_BORDERLESS = 1
const SCREEN_MODE_WINDOWED = 2