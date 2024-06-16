package pkg

import (
	"sync"

	"github.com/eddoog/be-capstone/models"
)

var TANJUNG_PRIOK = []string{
	"Ancol",
	"Cilincing",
	"Kalibaru",
	"Kamal Muara",
	"Kapuk Muara",
	"Kebon Bawang",
	"Kelapa Gading Barat",
	"Kelapa Gading Timur",
	"Koja",
	"Lagoa",
	"Marunda",
	"Pademangan Barat",
	"Pademangan Timur",
	"Papanggo",
	"Pegangsaan Dua",
	"Pejagalan",
	"Penjaringan",
	"Pluit",
	"Rawa Badak Selatan",
	"Rawa Badak Utara",
	"Rorotan",
	"Semper Barat",
	"Semper Timur",
	"Sukapura",
	"Sungai Bambu",
	"Sunter Agung",
	"Sunter Jaya",
	"Tanjung Priok",
	"Tugu Selatan",
	"Tugu Utara",
	"Warakas",
}

var KEMAYORAN = []string{
	"Cempaka Putih Barat",
	"Cempaka Putih Timur",
	"Rawasari",
	"Cideng",
	"Duri Pulo",
	"Gambir",
	"Kebon Kelapa",
	"Petojo Selatan",
	"Petojo Utara",
	"Galur",
	"Johar Baru",
	"Kampung Rawa",
	"Tanah Tinggi",
	"Cempaka Baru",
	"Gunung Sahari Selatan",
	"Harapan Mulya",
	"Kebon Kosong",
	"Kemayoran",
	"Serdang",
	"Sumur Batu",
	"Utan Panjang",
	"Cikini",
	"Gondangdia",
	"Kebon Sirih",
	"Menteng",
	"Pegangsaan",
	"Gunung Sahari Utara",
	"Karang Anyar",
	"Kartini",
	"Mangga Dua Selatan",
	"Pasar Baru",
	"Bungur",
	"Kenari",
	"Kramat",
	"Kwitang",
	"Paseban",
	"Senen",
	"Bendungan Hilir",
	"Gelora",
	"Kampung Bali",
	"Karet Tengsin",
	"Kebon Kacang",
	"Kebon Melati",
	"Petamburan",
	"Cengkareng Barat",
	"Cengkareng Timur",
	"Duri Kosambi",
	"Kapuk",
	"Kedaung Kali Angke",
	"Rawa Buaya",
	"Grogol",
	"Jelambar",
	"Jelambar Baru",
	"Tanjung Duren Selatan",
	"Tanjung Duren Utara",
	"Tomang",
	"Wijaya Kusuma",
	"Kalideres",
	"Kamal",
	"Pegadungan",
	"Semanan",
	"Tegal Alur",
	"Duri Kepa",
	"Kebon Jeruk",
	"Kedoya Selatan",
	"Kedoya Utara",
	"Kelapa Dua",
	"Sukabumi Selatan",
	"Sukabumi Utara",
	"Joglo",
	"Kembangan Selatan",
	"Kembangan Utara",
	"Meruya Selatan",
	"Meruya Utara",
	"Srengseng",
	"Jatipulo",
	"Kemanggisan",
	"Kota Bambu Selatan",
	"Kota Bambu Utara",
	"Palmerah",
	"Slipi",
	"Glodok",
	"Keagungan",
	"Krukut",
	"Mangga Besar",
	"Maphar",
	"Pinangsia",
	"Taman Sari",
	"Tangki",
	"Angke",
	"Duri Selatan",
	"Duri Utara",
	"Jembatan Besi",
	"Jembatan Lima",
	"Kali Anyar",
	"Krendang",
	"Pekojan",
	"Roa Malaka",
	"Tambora",
	"Tanah Sereal",
}

var HALIM = []string{
	"Balekambang",
	"Bali Mester",
	"Bambu Apus",
	"Baru",
	"Batu Ampar",
	"Bidara Cina",
	"Cakung Barat",
	"Cakung Timur",
	"Cawang",
	"Ceger",
	"Cibubur",
	"Cijantung",
	"Cilangkap",
	"Cililitan",
	"Cipayung",
	"Cipinang",
	"Cipinang Besar Selatan",
	"Cipinang Besar Utara",
	"Cipinang Cempedak",
	"Cipinang Melayu",
	"Cipinang Muara",
	"Ciracas",
	"Dukuh",
	"Duren Sawit",
	"Gedong",
	"Halim Perdana Kusuma",
	"Jati",
	"Jatinegara",
	"Jatinegara Kaum",
	"Kalisari",
	"Kampung Melayu",
	"Kayu Manis",
	"Kayu Putih",
	"Kebon Manggis",
	"Kebon Pala",
	"Kelapa Dua Wetan",
	"Klender",
	"Kramat Jati",
	"Lubang Buaya",
	"Makasar",
	"Malaka Jaya",
	"Malaka Sari",
	"Munjul",
	"Pal Meriam",
	"Pekayon",
	"Penggilingan",
	"Pinang Ranti",
	"Pisangan Baru",
	"Pisangan Timur",
	"Pondok Bambu",
	"Pondok Kelapa",
	"Pondok Kopi",
	"Pondok Ranggon",
	"Pulo Gadung",
	"Pulo Gebang",
	"Rambutan",
	"Rawa Bunga",
	"Rawa Terate",
	"Rawamangun",
	"Setu",
	"Susukan",
	"Tengah",
	"Ujung Menteng",
	"Utan Kayu Selatan",
	"Utan Kayu Utara",
	"Bangka",
	"Bintaro",
	"Bukit Duri",
	"Ciganjur",
	"Cikoko",
	"Cilandak Barat",
	"Cilandak Timur",
	"Cipedak",
	"Cipete Selatan",
	"Cipete Utara",
	"Cipulir",
	"Duren Tiga",
	"Gandaria Selatan",
	"Gandaria Utara",
	"Grogol Selatan",
	"Grogol Utara",
	"Guntur",
	"Gunung",
	"Jagakarsa",
	"Jati Padang",
	"Kalibata",
	"Karet",
	"Karet Kuningan",
	"Karet Semanggi",
	"Kebagusan",
	"Kebayoran Lama Selatan",
	"Kebayoran Lama Utara",
	"Kebon Baru",
	"Kramat Pela",
	"Kuningan Barat",
	"Kuningan Timur",
	"Lebak Bulus",
	"Lenteng Agung",
	"Mampang Prapatan",
	"Manggarai",
	"Manggarai Selatan",
	"Melawai",
	"Menteng Atas",
	"Menteng Dalam",
	"Pancoran",
	"Pasar Manggis",
	"Pasar Minggu",
	"Pejaten Barat",
	"Pejaten Timur",
	"Pela Mampang",
	"Pengadegan",
	"Pesanggrahan",
	"Petogogan",
	"Petukangan Selatan",
	"Petukangan Utara",
	"Pondok Labu",
	"Pondok Pinang",
	"Pulo",
	"Ragunan",
	"Rawa Barat",
	"Rawajati",
	"Selong",
	"Senayan",
	"Setiabudi",
	"Srengseng Sawah",
	"Tanjung Barat",
	"Tebet Barat",
	"Tebet Timur",
	"Tegal Parang",
	"Ulujami",
}

var (
	placesInstances *models.Places
	once            sync.Once
)

func InitializePlaces() {
	placesInstances = &models.Places{
		Halim: []models.Location{
			{LocationName: "Balekambang", Latitude: -6.278393599999999, Longitude: 106.8496169},
			{LocationName: "Bali Mester", Latitude: -6.2191246, Longitude: 106.8657771},
			{LocationName: "Bambu Apus", Latitude: -6.3092491, Longitude: 106.9033981},
			{LocationName: "Baru", Latitude: -6.326359399999999, Longitude: 106.8502879},
			{LocationName: "Batu Ampar", Latitude: -6.2748193, Longitude: 106.8620891},
			{LocationName: "Bidara Cina", Latitude: -6.2299555, Longitude: 106.8665147},
			{LocationName: "Cakung Barat", Latitude: -6.174699500000001, Longitude: 106.9363084},
			{LocationName: "Cakung Timur", Latitude: -6.177334399999999, Longitude: 106.9535691},
			{LocationName: "Cawang", Latitude: -6.2490679, Longitude: 106.8679899},
			{LocationName: "Ceger", Latitude: -6.304526000000001, Longitude: 106.8915948},
			{LocationName: "Cibubur", Latitude: -6.356087199999999, Longitude: 106.879792},
			{LocationName: "Cijantung", Latitude: -6.322277, Longitude: 106.8576636},
			{LocationName: "Cilangkap", Latitude: -6.334631000000001, Longitude: 106.9093},
			{LocationName: "Cililitan", Latitude: -6.263268399999999, Longitude: 106.8635643},
			{LocationName: "Cipayung", Latitude: -6.3153539, Longitude: 106.8928693},
			{LocationName: "Cipinang", Latitude: -6.207429299999999, Longitude: 106.8915948},
			{LocationName: "Cipinang Besar Selatan", Latitude: -6.2295694, Longitude: 106.8783167},
			{LocationName: "Cipinang Besar Utara", Latitude: -6.2193534, Longitude: 106.8783167},
			{LocationName: "Cipinang Cempedak", Latitude: -6.2386548, Longitude: 106.8738909},
			{LocationName: "Cipinang Melayu", Latitude: -6.2528204, Longitude: 106.9093},
			{LocationName: "Cipinang Muara", Latitude: -6.2176426, Longitude: 106.8915948},
			{LocationName: "Ciracas", Latitude: -6.308751699999999, Longitude: 106.8749344},
			{LocationName: "Dukuh", Latitude: -6.2960156, Longitude: 106.8783167},
			{LocationName: "Duren Sawit", Latitude: -6.2321915, Longitude: 106.9152021},
			{LocationName: "Gedong", Latitude: -6.3003881, Longitude: 106.8620891},
			{LocationName: "Halim Perdana Kusuma", Latitude: -6.246114200000001, Longitude: 106.8864928},
			{LocationName: "Jati", Latitude: -6.2086985, Longitude: 106.9461257},
			{LocationName: "Jatinegara", Latitude: -6.2284425, Longitude: 106.8851503},
			{LocationName: "Jatinegara Kaum", Latitude: -6.2032628, Longitude: 106.9019227},
			{LocationName: "Kalisari", Latitude: -6.3377275, Longitude: 106.8547133},
			{LocationName: "Kampung Melayu", Latitude: -6.215438199999999, Longitude: 106.8613515},
			{LocationName: "Kayu Manis", Latitude: -6.2026226, Longitude: 106.8628267},
			{LocationName: "Kayu Putih", Latitude: -6.1796396, Longitude: 106.8827427},
			{LocationName: "Kebon Manggis", Latitude: -6.209812599999999, Longitude: 106.8576636},
			{LocationName: "Kebon Pala", Latitude: -6.2551168, Longitude: 106.8783167},
			{LocationName: "Kelapa Dua Wetan", Latitude: -6.3456529, Longitude: 106.8856933},
			{LocationName: "Klender", Latitude: -6.2172566, Longitude: 106.9033981},
			{LocationName: "Kramat Jati", Latitude: -6.263320299999999, Longitude: 106.8656951},
			{LocationName: "Lubang Buaya", Latitude: -6.2939072, Longitude: 106.9033981},
			{LocationName: "Makasar", Latitude: -6.2652116, Longitude: 106.8758468},
			{LocationName: "Malaka Jaya", Latitude: -6.2226246, Longitude: 106.9343848},
			{LocationName: "Malaka Sari", Latitude: -6.2228182, Longitude: 106.9284822},
			{LocationName: "Munjul", Latitude: -6.3503773, Longitude: 106.8974964},
			{LocationName: "Pal Meriam", Latitude: -6.199599099999999, Longitude: 106.8576636},
			{LocationName: "Pekayon", Latitude: -6.3464395, Longitude: 106.8620891},
			{LocationName: "Penggilingan", Latitude: -6.2060792, Longitude: 106.9329091},
			{LocationName: "Pinang Ranti", Latitude: -6.290706699999999, Longitude: 106.884218},
			{LocationName: "Pisangan Baru", Latitude: -6.2119814, Longitude: 106.8694651},
			{LocationName: "Pisangan Timur", Latitude: -6.2078145, Longitude: 106.879792},
			{LocationName: "Pondok Bambu", Latitude: -6.2325788, Longitude: 106.9033981},
			{LocationName: "Pondok Kelapa", Latitude: -6.2393686, Longitude: 106.9299579},
			{LocationName: "Pondok Kopi", Latitude: -6.2248872, Longitude: 106.9432389},
			{LocationName: "Pondok Ranggon", Latitude: -6.3551012, Longitude: 106.9093},
			{LocationName: "Pulo Gadung", Latitude: -6.1936067, Longitude: 106.8912095},
			{LocationName: "Pulo Gebang", Latitude: -6.2079551, Longitude: 106.9535691},
			{LocationName: "Rambutan", Latitude: -6.3051125, Longitude: 106.8738909},
			{LocationName: "Rawa Bunga", Latitude: -6.2195462, Longitude: 106.8724156},
			{LocationName: "Rawa Terate", Latitude: -6.1860472, Longitude: 106.9211043},
			{LocationName: "Rawamangun", Latitude: -6.1976024, Longitude: 106.879792},
			{LocationName: "Setu", Latitude: -6.312742099999999, Longitude: 106.9137265},
			{LocationName: "Susukan", Latitude: -6.315538999999999, Longitude: 106.8679899},
			{LocationName: "Tengah", Latitude: -6.290004499999999, Longitude: 106.8672736},
			{LocationName: "Ujung Menteng", Latitude: -6.1899947, Longitude: 106.9565206},
			{LocationName: "Utan Kayu Selatan", Latitude: -6.2043213, Longitude: 106.8694651},
			{LocationName: "Utan Kayu Utara", Latitude: -6.196758099999999, Longitude: 106.8665147},
			{LocationName: "Bangka", Latitude: -6.2608384, Longitude: 106.8207875},
			{LocationName: "Bintaro", Latitude: -6.2678825, Longitude: 106.7617984},
			{LocationName: "Bukit Duri", Latitude: -6.222582, Longitude: 106.8576636},
			{LocationName: "Ciganjur", Latitude: -6.3379655, Longitude: 106.8089885},
			{LocationName: "Cikoko", Latitude: -6.245055799999999, Longitude: 106.8539757},
			{LocationName: "Cilandak Barat", Latitude: -6.2922975, Longitude: 106.79719},
			{LocationName: "Cilandak Timur", Latitude: -6.2791676, Longitude: 106.8152966},
			{LocationName: "Cipedak", Latitude: -6.3561798, Longitude: 106.8001396},
			{LocationName: "Cipete Selatan", Latitude: -6.2716448, Longitude: 106.8030892},
			{LocationName: "Cipete Utara", Latitude: -6.2596634, Longitude: 106.8080833},
			{LocationName: "Cipulir", Latitude: -6.23682, Longitude: 106.773595},
			{LocationName: "Duren Tiga", Latitude: -6.2553397, Longitude: 106.8325872},
			{LocationName: "Gandaria Selatan", Latitude: -6.2706081, Longitude: 106.7957153},
			{LocationName: "Gandaria Utara", Latitude: -6.258017700000001, Longitude: 106.7898163},
			{LocationName: "Grogol Selatan", Latitude: -6.2300428, Longitude: 106.7798811},
			{LocationName: "Grogol Utara", Latitude: -6.2159946, Longitude: 106.7853922},
			{LocationName: "Guntur", Latitude: -6.2080273, Longitude: 106.8340622},
			{LocationName: "Gunung", Latitude: -6.2349179, Longitude: 106.7927658},
			{LocationName: "Jagakarsa", Latitude: -6.3360601, Longitude: 106.8256645},
			{LocationName: "Jati Padang", Latitude: -6.286017699999999, Longitude: 106.8325872},
			{LocationName: "Kalibata", Latitude: -6.264140599999999, Longitude: 106.8370122},
			{LocationName: "Karet", Latitude: -6.2134231, Longitude: 106.8252123},
			{LocationName: "Karet Kuningan", Latitude: -6.2197608, Longitude: 106.8266873},
			{LocationName: "Karet Semanggi", Latitude: -6.2213742, Longitude: 106.8163628},
			{LocationName: "Kebagusan", Latitude: -6.3115948, Longitude: 106.8325872},
			{LocationName: "Kebayoran Lama Selatan", Latitude: -6.2570768, Longitude: 106.7794935},
			{LocationName: "Kebayoran Lama Utara", Latitude: -6.245621499999999, Longitude: 106.7780189},
			{LocationName: "Kebon Baru", Latitude: -6.2327029, Longitude: 106.8606139},
			{LocationName: "Kramat Pela", Latitude: -6.2451405, Longitude: 106.7927658},
			{LocationName: "Kuningan Barat", Latitude: -6.2365111, Longitude: 106.8222625},
			{LocationName: "Kuningan Timur", Latitude: -6.2299792, Longitude: 106.8266873},
			{LocationName: "Lebak Bulus", Latitude: -6.3031123, Longitude: 106.7794935},
			{LocationName: "Lenteng Agung", Latitude: -6.325502999999999, Longitude: 106.8341101},
			{LocationName: "Mampang Prapatan", Latitude: -6.2481661, Longitude: 106.8262565},
			{LocationName: "Manggarai", Latitude: -6.212558599999999, Longitude: 106.851763},
			{LocationName: "Manggarai Selatan", Latitude: -6.2203167, Longitude: 106.8488128},
			{LocationName: "Melawai", Latitude: -6.2448517, Longitude: 106.8016144},
			{LocationName: "Menteng Atas", Latitude: -6.218051099999999, Longitude: 106.8399623},
			{LocationName: "Menteng Dalam", Latitude: -6.2308233, Longitude: 106.8399623},
			{LocationName: "Pancoran", Latitude: -6.2554837, Longitude: 106.8429756},
			{LocationName: "Pasar Manggis", Latitude: -6.2090339, Longitude: 106.8415828},
			{LocationName: "Pasar Minggu", Latitude: -6.289610499999999, Longitude: 106.8399623},
			{LocationName: "Pejaten Barat", Latitude: -6.270482899999999, Longitude: 106.8384873},
			{LocationName: "Pejaten Timur", Latitude: -6.270095, Longitude: 106.8502879},
			{LocationName: "Pela Mampang", Latitude: -6.250807399999999, Longitude: 106.8148879},
			{LocationName: "Pengadegan", Latitude: -6.2482256, Longitude: 106.8547133},
			{LocationName: "Pesanggrahan", Latitude: -6.2365273, Longitude: 106.751266},
			{LocationName: "Petogogan", Latitude: -6.242007099999999, Longitude: 106.8104633},
			{LocationName: "Petukangan Selatan", Latitude: -6.2425078, Longitude: 106.7559004},
			{LocationName: "Petukangan Utara", Latitude: -6.2273641, Longitude: 106.7500025},
			{LocationName: "Pondok Labu", Latitude: -6.3076465, Longitude: 106.79719},
			{LocationName: "Pondok Pinang", Latitude: -6.2775326, Longitude: 106.7794935},
			{LocationName: "Pulo", Latitude: -6.255172, Longitude: 106.7986648},
			{LocationName: "Ragunan", Latitude: -6.301752, Longitude: 106.8207875},
			{LocationName: "Rawa Barat", Latitude: -6.2368, Longitude: 106.813413},
			{LocationName: "Rawajati", Latitude: -6.258447599999999, Longitude: 106.8547133},
			{LocationName: "Selong", Latitude: -6.2358591, Longitude: 106.8030892},
			{LocationName: "Senayan", Latitude: -6.2266758, Longitude: 106.8104633},
			{LocationName: "Setiabudi", Latitude: -6.2136327, Longitude: 106.8264328},
			{LocationName: "Srengseng Sawah", Latitude: -6.284691899999999, Longitude: 106.8189778},
			{LocationName: "Tanjung Barat", Latitude: -6.3052091, Longitude: 106.8483604},
			{LocationName: "Tebet Barat", Latitude: -6.2356436, Longitude: 106.8488128},
			{LocationName: "Tebet Timur", Latitude: -6.2328959, Longitude: 106.8547133},
			{LocationName: "Tegal Parang", Latitude: -6.249095199999999, Longitude: 106.8281623},
			{LocationName: "Ulujami", Latitude: -6.240989799999999, Longitude: 106.763273},
		},
		TanjungPriok: []models.Location{
			{LocationName: "Ancol", Latitude: -6.132975099999999, Longitude: 106.8266873},
			{LocationName: "Cilincing", Latitude: -6.1101657, Longitude: 106.9448985},
			{LocationName: "Kalibaru", Latitude: -6.104635, Longitude: 106.9152021},
			{LocationName: "Kamal Muara", Latitude: -6.1127546, Longitude: 106.7411559},
			{LocationName: "Kapuk Muara", Latitude: -6.1164245, Longitude: 106.7558436},
			{LocationName: "Kebon Bawang", Latitude: -6.1194606, Longitude: 106.8901194},
			{LocationName: "Kelapa Gading Barat", Latitude: -6.1163473, Longitude: 106.7022381},
			{LocationName: "Kelapa Gading Timur", Latitude: -6.170309, Longitude: 106.9033058},
			{LocationName: "Koja", Latitude: -6.115910299999999, Longitude: 106.9081259},
			{LocationName: "Lagoa", Latitude: -6.1150195, Longitude: 106.9093},
			{LocationName: "Marunda", Latitude: -6.1133063, Longitude: 106.9624239},
			{LocationName: "Pademangan Barat", Latitude: -6.133918599999999, Longitude: 106.8370122},
			{LocationName: "Pademangan Timur", Latitude: -6.147520399999999, Longitude: 106.8502879},
			{LocationName: "Papanggo", Latitude: -6.1265467, Longitude: 106.8679899},
			{LocationName: "Pegangsaan Dua", Latitude: -6.1658278, Longitude: 106.9152021},
			{LocationName: "Pejagalan", Latitude: -6.138162599999999, Longitude: 106.7822166},
			{LocationName: "Penjaringan", Latitude: -6.116455699999999, Longitude: 106.7557702},
			{LocationName: "Pluit", Latitude: -6.1128512, Longitude: 106.7843739},
			{LocationName: "Rawa Badak Selatan", Latitude: -6.131923, Longitude: 106.8989718},
			{LocationName: "Rawa Badak Utara", Latitude: -6.1204977, Longitude: 106.8974964},
			{LocationName: "Rorotan", Latitude: -6.1467296, Longitude: 106.9535691},
			{LocationName: "Semper Barat", Latitude: -6.1259668, Longitude: 106.925531},
			{LocationName: "Semper Timur", Latitude: -6.1193561, Longitude: 106.9329091},
			{LocationName: "Sukapura", Latitude: -6.1474953, Longitude: 106.9299579},
			{LocationName: "Sungai Bambu", Latitude: -6.1361763, Longitude: 106.8856933},
			{LocationName: "Sunter Agung", Latitude: -6.136937199999999, Longitude: 106.8620891},
			{LocationName: "Sunter Jaya", Latitude: -6.1502614, Longitude: 106.8795785},
			{LocationName: "Tanjung Priok", Latitude: -6.1382799, Longitude: 106.8665166},
			{LocationName: "Tugu Selatan", Latitude: -6.1354142, Longitude: 106.9093},
			{LocationName: "Tugu Utara", Latitude: -6.1252159, Longitude: 106.9093},
			{LocationName: "Warakas", Latitude: -6.1223898, Longitude: 106.8783167},
		},
		Kemayoran: []models.Location{
			{LocationName: "Cempaka Putih Barat", Latitude: -6.1815383, Longitude: 106.8635643},
			{LocationName: "Cempaka Putih Timur", Latitude: -6.1761468, Longitude: 106.8724156},
			{LocationName: "Rawasari", Latitude: -6.1890998, Longitude: 106.8665147},
			{LocationName: "Cideng", Latitude: -6.1705904, Longitude: 106.8075136},
			{LocationName: "Duri Pulo", Latitude: -6.163123499999999, Longitude: 106.8016144},
			{LocationName: "Gambir", Latitude: -6.1673595, Longitude: 106.8209277},
			{LocationName: "Kebon Kelapa", Latitude: -6.1649144, Longitude: 106.8252123},
			{LocationName: "Petojo Selatan", Latitude: -6.1754096, Longitude: 106.8163628},
			{LocationName: "Petojo Utara", Latitude: -6.163971600000001, Longitude: 106.8148879},
			{LocationName: "Galur", Latitude: -6.1760588, Longitude: 106.8554508},
			{LocationName: "Johar Baru", Latitude: -6.1830538, Longitude: 106.8561884},
			{LocationName: "Kampung Rawa", Latitude: -6.1798872, Longitude: 106.8554508},
			{LocationName: "Tanah Tinggi", Latitude: -6.179464299999999, Longitude: 106.8488128},
			{LocationName: "Cempaka Baru", Latitude: -6.1687781, Longitude: 106.8635643},
			{LocationName: "Gunung Sahari Selatan", Latitude: -6.1579159, Longitude: 106.8443875},
			{LocationName: "Harapan Mulya", Latitude: -6.1716166, Longitude: 106.8547133},
			{LocationName: "Kebon Kosong", Latitude: -6.1563067, Longitude: 106.8547133},
			{LocationName: "Kemayoran", Latitude: -6.1594869, Longitude: 106.8541776},
			{LocationName: "Serdang", Latitude: -6.1562991, Longitude: 106.8615947},
			{LocationName: "Sumur Batu", Latitude: -6.163388299999999, Longitude: 106.8724156},
			{LocationName: "Utan Panjang", Latitude: -6.1658988, Longitude: 106.8539757},
			{LocationName: "Cikini", Latitude: -6.192514999999999, Longitude: 106.8399623},
			{LocationName: "Gondangdia", Latitude: -6.192897899999999, Longitude: 106.8281623},
			{LocationName: "Kebon Sirih", Latitude: -6.1851432, Longitude: 106.8311122},
			{LocationName: "Menteng", Latitude: -6.1937297, Longitude: 106.8392111},
			{LocationName: "Pegangsaan", Latitude: -6.2024403, Longitude: 106.8488128},
			{LocationName: "Gunung Sahari Utara", Latitude: -6.149223999999999, Longitude: 106.8370122},
			{LocationName: "Karang Anyar", Latitude: -6.153950399999999, Longitude: 106.8288997},
			{LocationName: "Kartini", Latitude: -6.152532, Longitude: 106.8333247},
			{LocationName: "Mangga Dua Selatan", Latitude: -6.1418556, Longitude: 106.8281623},
			{LocationName: "Pasar Baru", Latitude: -6.1671806, Longitude: 106.8340622},
			{LocationName: "Bungur", Latitude: -6.171807599999999, Longitude: 106.8488128},
			{LocationName: "Kenari", Latitude: -6.1929378, Longitude: 106.8466001},
			{LocationName: "Kramat", Latitude: -6.184002899999999, Longitude: 106.8466001},
			{LocationName: "Kwitang", Latitude: -6.1829179, Longitude: 106.8406998},
			{LocationName: "Paseban", Latitude: -6.192131799999999, Longitude: 106.851763},
			{LocationName: "Senen", Latitude: -6.195351899999999, Longitude: 106.8498184},
			{LocationName: "Bendungan Hilir", Latitude: -6.2088902, Longitude: 106.8075136},
			{LocationName: "Gelora", Latitude: -6.2154197, Longitude: 106.8030892},
			{LocationName: "Kampung Bali", Latitude: -6.1856209, Longitude: 106.8163628},
			{LocationName: "Karet Tengsin", Latitude: -6.2048197, Longitude: 106.8148879},
			{LocationName: "Kebon Kacang", Latitude: -6.1907272, Longitude: 106.8163628},
			{LocationName: "Kebon Melati", Latitude: -6.1983875, Longitude: 106.8163628},
			{LocationName: "Petamburan", Latitude: -6.1961208, Longitude: 106.8075136},
			{LocationName: "Cengkareng Barat", Latitude: -6.1361871, Longitude: 106.7264125},
			{LocationName: "Cengkareng Timur", Latitude: -6.1370178, Longitude: 106.7330991},
			{LocationName: "Duri Kosambi", Latitude: -6.1721148, Longitude: 106.7205154},
			{LocationName: "Kapuk", Latitude: -6.1403479, Longitude: 106.7559004},
			{LocationName: "Kedaung Kali Angke", Latitude: -6.150556, Longitude: 106.7559004},
			{LocationName: "Rawa Buaya", Latitude: -6.1613339, Longitude: 106.7382072},
			{LocationName: "Grogol", Latitude: -6.161697999999999, Longitude: 106.7846412},
			{LocationName: "Jelambar", Latitude: -6.159817599999999, Longitude: 106.7853922},
			{LocationName: "Jelambar Baru", Latitude: -6.1488247, Longitude: 106.7847391},
			{LocationName: "Tanjung Duren Selatan", Latitude: -6.17882, Longitude: 106.7898163},
			{LocationName: "Tanjung Duren Utara", Latitude: -6.1713515, Longitude: 106.7839175},
			{LocationName: "Tomang", Latitude: -6.1747751, Longitude: 106.7999513},
			{LocationName: "Wijaya Kusuma", Latitude: -6.153768899999999, Longitude: 106.7750696},
			{LocationName: "Kalideres", Latitude: -6.1568457, Longitude: 106.7038324},
			{LocationName: "Kamal", Latitude: -6.1014059, Longitude: 106.6969285},
			{LocationName: "Pegadungan", Latitude: -6.1318356, Longitude: 106.702825},
			{LocationName: "Semanan", Latitude: -6.1649271, Longitude: 106.7057733},
			{LocationName: "Tegal Alur", Latitude: -6.118538699999999, Longitude: 106.7190298},
			{LocationName: "Duri Kepa", Latitude: -6.169184599999999, Longitude: 106.7748596},
			{LocationName: "Kebon Jeruk", Latitude: -6.184046800000001, Longitude: 106.7680752},
			{LocationName: "Kedoya Selatan", Latitude: -6.1810006, Longitude: 106.7617984},
			{LocationName: "Kedoya Utara", Latitude: -6.1707876, Longitude: 106.7617984},
			{LocationName: "Kelapa Dua", Latitude: -6.2101341, Longitude: 106.7691712},
			{LocationName: "Sukabumi Selatan", Latitude: -6.221487499999999, Longitude: 106.773595},
			{LocationName: "Sukabumi Utara", Latitude: -6.2098473, Longitude: 106.7780189},
			{LocationName: "Joglo", Latitude: -6.218219100000001, Longitude: 106.7453307},
			{LocationName: "Kembangan Selatan", Latitude: -6.186754, Longitude: 106.7430842},
			{LocationName: "Kembangan Utara", Latitude: -6.1713568, Longitude: 106.7441048},
			{LocationName: "Meruya Selatan", Latitude: -6.2087246, Longitude: 106.7337841},
			{LocationName: "Meruya Utara", Latitude: -6.1970854, Longitude: 106.7382072},
			{LocationName: "Srengseng", Latitude: -6.204272599999999, Longitude: 106.7529514},
			{LocationName: "Jatipulo", Latitude: -6.1783437, Longitude: 106.804564},
			{LocationName: "Kemanggisan", Latitude: -6.190262199999999, Longitude: 106.791291},
			{LocationName: "Kota Bambu Selatan", Latitude: -6.187347, Longitude: 106.8014743},
			{LocationName: "Kota Bambu Utara", Latitude: -6.1836404, Longitude: 106.7986648},
			{LocationName: "Palmerah", Latitude: -6.1942907, Longitude: 106.793333},
			{LocationName: "Slipi", Latitude: -6.1937584, Longitude: 106.8016144},
			{LocationName: "Glodok", Latitude: -6.144881199999999, Longitude: 106.813413},
			{LocationName: "Keagungan", Latitude: -6.1505982, Longitude: 106.8141505},
			{LocationName: "Krukut", Latitude: -6.157639199999999, Longitude: 106.813413},
			{LocationName: "Mangga Besar", Latitude: -6.145353000000001, Longitude: 106.8185752},
			{LocationName: "Maphar", Latitude: -6.1567398, Longitude: 106.821525},
			{LocationName: "Pinangsia", Latitude: -6.134676799999999, Longitude: 106.813413},
			{LocationName: "Taman Sari", Latitude: -6.1500077, Longitude: 106.816861},
			{LocationName: "Tangki", Latitude: -6.1471478, Longitude: 106.8222625},
			{LocationName: "Angke", Latitude: -6.1454501, Longitude: 106.7957153},
			{LocationName: "Duri Selatan", Latitude: -6.1585385, Longitude: 106.8053014},
			{LocationName: "Duri Utara", Latitude: -6.1547106, Longitude: 106.8053014},
			{LocationName: "Jembatan Besi", Latitude: -6.153010300000001, Longitude: 106.7986648},
			{LocationName: "Jembatan Lima", Latitude: -6.1451657, Longitude: 106.804564},
			{LocationName: "Kali Anyar", Latitude: -6.1574525, Longitude: 106.7994022},
			{LocationName: "Krendang", Latitude: -6.1496546, Longitude: 106.8038266},
			{LocationName: "Pekojan", Latitude: -6.1375119, Longitude: 106.804564},
			{LocationName: "Roa Malaka", Latitude: -6.1367084, Longitude: 106.8097259},
			{LocationName: "Tambora", Latitude: -6.1477051, Longitude: 106.8072461},
			{LocationName: "Tanah Sereal", Latitude: -6.1532923, Longitude: 106.8097259},
		},
	}
}

func GetPlaces() *models.Places {
	once.Do(InitializePlaces)

	return placesInstances
}
