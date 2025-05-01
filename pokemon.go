package faker

var pokemonJapanese = []string{
	"フシギダネ", "フシギソウ", "フシギバナ", "ヒトカゲ", "リザード", "リザードン", "ゼニガメ", "カメール", "カメックス", "キャタピー",
	"トランセル", "バタフリー", "ビードル", "コクーン", "スピアー", "ポッポ", "ピジョン", "ピジョット", "コラッタ", "ラッタ",
	"オニスズメ", "オニドリル", "アーボ", "アーボック", "ピカチュウ", "ライチュウ", "サンド", "サンドパン", "ニドラン♀（メス）", "ニドリーナ",
	"ニドクイン", "ニドラン♂（オス）", "ニドリーノ", "ニドキング", "ピッピ", "ピクシー", "ロコン", "キュウコン", "プリン", "プクリン",
	"ズバット", "ゴルバット", "ナゾノクサ", "クサイハナ", "ラフレシア", "パラセクト", "パラセクト", "コンパン", "モルフォン", "ディグダ",
	"ダグトリオ", "ニャース", "ペルシアン", "コダック", "ゴルダ", "マチス", "ゴリラ", "ガーディ", "ウインディ", "ニョロモ",
	"ニョロゾ", "ニョロボン", "ケーシィ", "ユンゲラー", "フーディン", "ワンリキー", "ゴーリキー", "カイリキー", "マダツボミ", "ウツドン",
	"ウツボット", "メノクラゲ", "ドククラゲ", "イワーク", "ゴローン", "ゴローニャ", "ポニータ", "ギャロップ", "ヤドン", "ヤドラン",
	"コイル", "レアコイル", "カモネギ", "ドードー", "ドードリオ", "パルシェン", "ジュゴン", "ベトベター", "ベトベトン", "シェルダー",
	"パルシェン", "ゴース", "ゲンガー", "ゲンガー", "イワーク", "スリープ", "スリーパー", "クラブ", "キングラー", "ビリリダマ",
	"マルマイン", "タマタマ", "ナッシー", "カラカラ", "ガラガラ", "サワムラー", "エビワラー", "ベロリンガ", "ドガース", "マタドガス",
	"サイホーン", "サイドン", "ラッキー", "モンジャラ", "ガルーラ", "タッツー", "シードラ", "コイキング", "ギャラドス", "ヒトデマン",
	"スターミー", "バリヤード", "ストライク", "ルージュラ", "エレブー", "ブーバー", "カイロス", "カイリュー", "コイキング", "ギャラドス",
	"ラプラス", "メタモン", "イーブイ", "シャワーズ", "サンダース", "ブースター", "ポリゴン", "オムナイト", "オムスター", "カブト",
	"カブトプス", "プテラ", "カビゴン", "フリーザ", "サンダー", "ファイヤー", "ミニリュウ", "ハクリュー", "カイリュー", "ミュウツー",
	"ミュウ", "チコリータ", "ベイリーフ", "メガニウム", "ヒノアラシ", "マリル", "バクフーン", "ワニノコ", "アリゲイツ", "オーダイル",
	"オタチ", "オオタチ", "ホーホー", "ヨルノズク", "レディバ", "レディアン", "イトマル", "アリアドス", "クロバット", "コイル",
	"ランターン", "ピチュー", "ピィ", "ププリン", "トゲピー", "トゲチック", "ネイティ", "ネイティオ", "メリープ", "モココ",
	"デンリュウ", "キレイハナ", "マリル", "マリルリ", "ウソッキー", "ニョロトノ", "ハネッコ", "ポポッコ", "ワタッコ", "エイプム",
	"ヒマナッツ", "キマワリ", "ヤンヤンマ", "ウパー", "ヌオー", "エーフィ", "ブラッキー", "ヤミカラス", "ヤドキング", "ムウマ",
	"アンノーン", "ソーナンス", "キリンリキ", "ドンファン", "グライガー", "ハガネール", "ブルー", "グランブル", "ハリーセン", "ツボツボ",
	"ヘラクロス", "ニューラ", "ヒマナッツ", "リングマ", "マグマッグ", "マグカルゴ", "ウリムー", "マンムー", "ホエルコ", "チョンチー",
	"オクタン", "デリバード", "マンタイン", "エアームド", "デルビル", "ヘルガー", "キングドラ", "ドンファン", "ドンファン", "ポリゴン2",
	"オドシシ", "ドーブル", "ヒノアラシ", "カポエラー", "ムチュール", "エレキッド", "ブビィ", "ミルタンク", "ハピナス", "ラッキー",
	"ヤングース", "ドククラゲ", "ギャラドス", "シャーク", "ホエルコ", "ホエルオー", "ドンメル", "バクーダ", "コータス", "バネブー",
	"ブーピッグ", "パッチール", "アントリット", "ビブラーバ", "フライゴン", "マクノシタ", "ハリテヤマ", "マリル", "ノズパス", "エネコ",
	"エネコロロ", "闇夜騎士", "クチート", "ココドラ", "ドサイドン", "ドサイドン", "メディタ", "メディタ", "ラクライ", "ライボルト",
	"プラスル", "マイナン", "バルビート", "イルミーゼ", "ロゼリア", "ゴニョニョ", "マルノーム", "ラブカス", "ヒノヤコマ", "ラティアス",
	"キャモメ", "テッカニン", "テッカニン", "ヌケニン", "ヤジロン", "ドサイドン", "アンモナイト", "アーマルド", "カブト", "アーマルド",
	"コイキング", "ミロカロス", "キャモメ", "キャモメ", "キャモメ", "ユレイドル", "ロゼリア", "ドンカラス", "ハスボー", "ロゼリア",
	"ドンカラス", "タネボー", "コノハナ", "ドンカラス", "スバメ", "オオスバメ", "キャモメ", "ペリッパー", "ラルトス", "キルリア",
	"サーナイト", "タツベイ", "コドラ", "ボーマンダ", "ダンバル", "メタグロス", "メタグロス", "レジロック", "レジアイス", "レジスチル",
	"ラティアス", "ラティオス", "カイオーガ", "グラードン", "レックウザ", "ジラーチ", "デオキシス", "デオキシス", "デオキシス", "デオキシス",
	"ナエトル", "ハヤシガメ", "ドダイトス", "ヒコザル", "ゴウカザル", "ゴウカザル", "ポッチャマ", "ポッタイシ", "エンペルト", "ムックル",
	"ムクバード", "ムクホーク", "ビッパ", "ビーダル", "コロボーシ", "コロトック", "コリンク", "ルクシオ", "レントラー", "スボミー",
	"ロズレイド", "ズガイドス", "ラムパルド", "タテトプス", "トリデプス", "ミノムッチ", "ミノマダム", "ミノマダム", "ミノマダム", "ガメノデス",
	"ミツハニー", "ビークイン", "パチリス", "ブイゼル", "フローゼル", "チェリン", "チェリム", "カラナクシ", "トリトドン", "カラナクシ",
	"トリトドン", "エテボース", "フワンテ", "フワライド", "ミミロル", "ミミロップ", "ムウマージ", "ドンカラス", "ニャルマー", "ブニャット",
	"スカタンク", "スカタンク", "ドータクン", "ドーミラー", "バリヤード", "ラッキー", "ペラップ", "ギラティナ", "フカマル", "ガバイト",
	"ガブリアス", "ゴンベ", "リオル", "ルカリオ", "ヒポポタス", "カバルドン", "スコルピ", "ドラピオン", "ヤングース", "ドククラゲ",
	"ウツボット", "コイキング", "ルンパッパ", "マンタイン", "ユキノオー", "ユキノオー", "マニューラ", "ジバコイル", "ベロベルト", "ドサイドン",
	"モジャン", "エレキブル", "ブーバー", "トゲキッス", "ヤンヤンマ", "リーフィア", "グレイシア", "グライオン", "マンムー", "ポリゴンZ",
	"エルレイド", "ダイノーズ", "ヨノワール", "ユキメノコ", "ロトム", "ギラティナ", "ギラティナ", "クレセリア", "マナフィ", "マナフィ",
	"ダークライ", "シェイミ", "シェイミ", "アルセウス",
}

var pokemonEnglish = []string{
	"Bulbasaur", "Ivysaur", "Venusaur", "Charmander", "Charmeleon", "Charizard", "Squirtle", "Wartortle", "Blastoise", "Caterpie",
	"Metapod", "Butterfree", "Weedle", "Kakuna", "Beedrill", "Pidgey", "Pidgeotto", "Pidgeot", "Rattata", "Raticate",
	"Spearow", "Fearow", "Ekans", "Arbok", "Pikachu", "Raichu", "Sandshrew", "Sandslash", "Nidoran♀", "Nidorina",
	"Nidoqueen", "Nidoran♂", "Nidorino", "Nidoking", "Clefairy", "Clefable", "Vulpix", "Ninetales", "Jigglypuff", "Wigglytuff",
	"Zubat", "Golbat", "Oddish", "Gloom", "Vileplume", "Paras", "Parasect", "Venonat", "Venomoth", "Diglett",
	"Dugtrio", "Meowth", "Persian", "Psyduck", "Golduck", "Mankey", "Primeape", "Growlithe", "Arcanine", "Poliwag",
	"Poliwhirl", "Poliwrath", "Abra", "Kadabra", "Alakazam", "Machop", "Machoke", "Machamp", "Bellsprout", "Weepinbell",
	"Victreebel", "Tentacool", "Tentacruel", "Geodude", "Graveler", "Golem", "Ponyta", "Rapidash", "Slowpoke", "Slowbro",
	"Magnemite", "Magneton", "Farfetch'd", "Doduo", "Dodrio", "Seel", "Dewgong", "Grimer", "Muk", "Shellder",
	"Cloyster", "Gastly", "Haunter", "Gengar", "Onix", "Drowzee", "Hypno", "Krabby", "Kingler", "Voltorb",
	"Electrode", "Exeggcute", "Exeggutor", "Cubone", "Marowak", "Hitmonlee", "Hitmonchan", "Lickitung", "Koffing", "Weezing",
	"Rhyhorn", "Rhydon", "Chansey", "Tangela", "Kangaskhan", "Horsea", "Seadra", "Goldeen", "Seaking", "Staryu",
	"Starmie", "Mr. Mime", "Scyther", "Jynx", "Electabuzz", "Magmar", "Pinsir", "Tauros", "Magikarp", "Gyarados",
	"Lapras", "Ditto", "Eevee", "Vaporeon", "Jolteon", "Flareon", "Porygon", "Omanyte", "Omastar", "Kabuto",
	"Kabutops", "Aerodactyl", "Snorlax", "Articuno", "Zapdos", "Moltres", "Dratini", "Dragonair", "Dragonite", "Mewtwo",
	"Mew", "Chikorita", "Bayleef", "Meganium", "Cyndaquil", "Quilava", "Typhlosion", "Totodile", "Croconaw", "Feraligatr",
	"Sentret", "Furret", "Hoothoot", "Noctowl", "Ledyba", "Ledian", "Spinarak", "Ariados", "Crobat", "Chinchou",
	"Lanturn", "Pichu", "Cleffa", "Igglybuff", "Togepi", "Togetic", "Natu", "Xatu", "Mareep", "Flaaffy",
	"Ampharos", "Bellossom", "Marill", "Azumarill", "Sudowoodo", "Politoed", "Hoppip", "Skiploom", "Jumpluff", "Aipom",
	"Sunkern", "Sunflora", "Yanma", "Wooper", "Quagsire", "Espeon", "Umbreon", "Murkrow", "Slowking", "Misdreavus",
	"Unown", "Wobbuffet", "Girafarig", "Dunsparce", "Gligar", "Steelix", "Snubbull", "Granbull", "Qwilfish", "Shuckle",
	"Heracross", "Sneasel", "Teddiursa", "Ursaring", "Slugma", "Magcargo", "Swinub", "Piloswine", "Corsola", "Remoraid",
	"Octillery", "Delibird", "Mantine", "Skarmory", "Houndour", "Houndoom", "Kingdra", "Phanpy", "Donphan", "Porygon2",
	"Stantler", "Smeargle", "Tyrogue", "Hitmontop", "Smoochum", "Elekid", "Magby", "Miltank", "Blissey", "Happiny",
	"Coragunk", "Toxicroak", "Carvanha", "Sharpedo", "Wailmer", "Wailord", "Numel", "Camerupt", "Torkoal", "Spoink",
	"Grumpig", "Spinda", "Trapinch", "Vibrava", "Flygon", "Makuhita", "Hariyama", "Azurill", "Nosepass", "Skitty",
	"Delcatty", "Sableye", "Mawile", "Aron", "Lairon", "Aggron", "Meditite", "Medicham", "Electrike", "Manectric",
	"Plusle", "Minun", "Volbeat", "Illumise", "Roselia", "Gulpin", "Swalot", "Luvdisc", "Solrock", "Lunatone",
	"Castform", "Nincada", "Ninjask", "Shedinja", "Baltoy", "Claydol", "Lileep", "Cradily", "Anorith", "Armaldo",
	"Feebas", "Milotic", "Castform (Snowy Form)", "Castform (Rainy Form)", "Castform (Sunny Form)", "Relicanth", "Lombre",
	"Ludicolo", "Lotad", "Lombre", "Ludicolo", "Seedot", "Nuzleaf", "Shiftry", "Taillow", "Swellow", "Wingull",
	"Pelipper", "Ralts", "Kirlia", "Gardevoir", "Bagon", "Shelgon", "Salamence", "Beldum", "Metang", "Metagross",
	"Regirock", "Regice", "Registeel", "Latias", "Latios", "Kyogre", "Groudon", "Rayquaza", "Jirachi", "Deoxys (Normal Form)",
	"Deoxys (Attack Form)", "Deoxys (Defense Form)", "Deoxys (Speed Form)", "Turtwig", "Grotle", "Torterra", "Chimchar",
	"Monferno", "Infernape", "Piplup", "Prinplup", "Empoleon", "Starly", "Staravia", "Staraptor", "Bidoof", "Bibarel",
	"Kricketot", "Kricketune", "Shinx", "Luxio", "Luxray", "Budew", "Roserade", "Cranidos", "Rampardos", "Shieldon",
	"Bastiodon", "Burmy", "Wormadam (Plant Cloak)", "Wormadam (Sandy Cloak)", "Wormadam (Trash Cloak)", "Mothim", "Combee",
	"Vespiquen", "Pachirisu", "Buizel", "Floatzel", "Cherubi", "Cherrim", "Shellos (West Sea)", "Gastrodon (West Sea)",
	"Shellos (East Sea)", "Gastrodon (East Sea)", "Ambipom", "Drifloon", "Drifblim", "Buneary", "Lopunny", "Mismagius",
	"Honchkrow", "Glameow", "Purugly", "Stunky", "Skuntank", "Bronzong", "Bonsly", "Mime Jr.", "Happiny", "Chatot",
	"Spiritomb", "Gible", "Gabite", "Garchomp", "Munchlax", "Riolu", "Lucario", "Hippopotas", "Hippowdon", "Skorupi",
	"Drapion", "Croagunk", "Toxicroak", "Carnivine", "Finneon", "Lumineon", "Mantyke", "Snover", "Abomasnow", "Weavile",
	"Magnezone", "Lickilicky", "Rhyperior", "Tangrowth", "Electivire", "Magmortar", "Togekiss", "Yanmega", "Leafeon",
	"Glaceon", "Gliscor", "Mamoswine", "Porygon-Z", "Gallade", "Probopass", "Dusknoir", "Froslass", "Rotom",
	"Giratina (Altered Form)", "Giratina (Origin Form)", "Cresselia", "Phione", "Manaphy", "Darkrai", "Shaymin (Land Form)",
	"Shaymin (Sky Form)", "Arceus",
}

// Pokemon is a faker struct for Pokemon
type Pokemon struct {
	Faker *Faker
}

// Japanese returns a fake Japanese pokemon name
func (p Pokemon) Japanese() string {
	return p.Faker.RandomStringElement(pokemonJapanese)
}

// English returns a fake English pokemon name
func (p Pokemon) English() string {
	return p.Faker.RandomStringElement(pokemonEnglish)
}
