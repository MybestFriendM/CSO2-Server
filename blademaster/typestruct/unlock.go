package typestruct

type UnlockItem struct {
	Itemid   uint32
	Seq      uint32
	CostType uint8
	Price    uint32
}

type UnlockKillNum struct {
	Unk00   uint32 //seq ?
	Itemid  uint32
	KillNum uint32
	Unk01   [6]byte //always 0
}

type UnlockSeq struct {
	Nums uint16
	IDs  []uint32
}

var (
	UnlockItemList = [...]uint32{
		1,
		9,
		10,
		24,
		11,
		60,
		31,
		17,
		28,
		59,
		53,
		26,
		25,
		63,
		20,
		7,
		62,
		5,
		44,
		16,
		12,
		22,
		52,
		70,
		71,
		77,
		85,
		48,
		29,
		32,
		37,
		106,
		112,
		111,
		110,
		105,
		114,
		107,
		109,
		74,
		75,
		78,
		82,
		91,
		95,
		96,
		100,
		104,
		109,
		108,
		110,
		66,
		67,
		120,
		121,
		124,
		122,
		123,
		125,
		126,
		129,
		130,
		131,
		133,
		132,
		135,
		138,
		143,
		144,
		145,
		155,
		156,
		151,
		152,
		153,
	}
	UnlockFullList = []UnlockItem{
		UnlockItem{
			UnlockItemList[0],
			0x0B,
			0x01,
			0x3E8,
		},
		UnlockItem{
			UnlockItemList[1],
			0x0C,
			0x01,
			0x5DC,
		},
		UnlockItem{
			UnlockItemList[2],
			0x0D,
			0x01,
			0x3E8,
		},
		UnlockItem{
			UnlockItemList[3],
			0x0E,
			0x01,
			0x5DC,
		},
		UnlockItem{
			UnlockItemList[4],
			0x0F,
			0x01,
			0x708,
		},
		UnlockItem{
			UnlockItemList[5],
			0x10,
			0x01,
			0xBB80,
		},
		UnlockItem{
			UnlockItemList[6],
			0x11,
			0x01,
			0x5DC0,
		},
		UnlockItem{
			UnlockItemList[7],
			0x12,
			0x01,
			0x708,
		},
		UnlockItem{
			UnlockItemList[8],
			0x13,
			0x01,
			0x1D4C,
		},
		UnlockItem{
			UnlockItemList[9],
			0x14,
			0x01,
			0x6160,
		},
		UnlockItem{
			UnlockItemList[10],
			0x15,
			0x01,
			0x7530,
		},
		UnlockItem{
			UnlockItemList[11],
			0x16,
			0x01,
			0xFA0,
		},
		UnlockItem{
			UnlockItemList[12],
			0x17,
			0x01,
			0x3A98,
		},
		UnlockItem{
			UnlockItemList[13],
			0x18,
			0x01,
			0x493E0,
		},
		UnlockItem{
			UnlockItemList[14],
			0x19,
			0x01,
			0xFA0,
		},
		UnlockItem{
			UnlockItemList[15],
			0x1A,
			0x01,
			0x3A98,
		},
		UnlockItem{
			UnlockItemList[16],
			0x1B,
			0x01,
			0x493E0,
		},
		UnlockItem{
			UnlockItemList[17],
			0x1C,
			0x01,
			0x708,
		},
		UnlockItem{
			UnlockItemList[18],
			0x1D,
			0x01,
			0x7530,
		},
		UnlockItem{
			UnlockItemList[19],
			0x1E,
			0x01,
			0x1388,
		},
		UnlockItem{
			UnlockItemList[20],
			0x1F,
			0x01,
			0x4E20,
		},
		UnlockItem{
			UnlockItemList[21],
			0x20,
			0x01,
			0x4E20,
		},
		UnlockItem{
			UnlockItemList[22],
			0x43,
			0x01,
			0x7530,
		},
		UnlockItem{
			UnlockItemList[23],
			0x57,
			0x01,
			0x7A120,
		},
		UnlockItem{
			UnlockItemList[24],
			0x58,
			0x01,
			0x7A120,
		},
		UnlockItem{
			UnlockItemList[25],
			0x59,
			0x01,
			0x190,
		},
		UnlockItem{
			UnlockItemList[26],
			0x81,
			0x01,
			0x370,
		},
		UnlockItem{
			UnlockItemList[27],
			0x90,
			0x01,
			0x7530,
		},
		UnlockItem{
			UnlockItemList[28],
			0x91,
			0x01,
			0xEA60,
		},
		UnlockItem{
			UnlockItemList[29],
			0x92,
			0x01,
			0x1E848,
		},
		UnlockItem{
			UnlockItemList[30],
			0x93,
			0x01,
			0x30D40,
		},
		UnlockItem{
			UnlockItemList[31],
			0xA8,
			0x00,
			0x28,
		},
		UnlockItem{
			UnlockItemList[32],
			0xA9,
			0x00,
			0x50,
		},
		UnlockItem{
			UnlockItemList[33],
			0xAA,
			0x00,
			0x28,
		},
		UnlockItem{
			UnlockItemList[34],
			0xAB,
			0x00,
			0x50,
		},
		UnlockItem{
			UnlockItemList[35],
			0xAC,
			0x00,
			0x28,
		},
		UnlockItem{
			UnlockItemList[36],
			0xAD,
			0x00,
			0x50,
		},
		UnlockItem{
			UnlockItemList[37],
			0xAE,
			0x00,
			0x28,
		},
		UnlockItem{
			UnlockItemList[38],
			0xAF,
			0x00,
			0x50,
		},
		UnlockItem{
			UnlockItemList[39],
			0xD7,
			0x01,
			0xC350,
		},
		UnlockItem{
			UnlockItemList[40],
			0xD8,
			0x01,
			0x17700,
		},
		UnlockItem{
			UnlockItemList[41],
			0xE8,
			0x01,
			0x11170,
		},
		UnlockItem{
			UnlockItemList[42],
			0xE9,
			0x01,
			0x1D4C0,
		},
		UnlockItem{
			UnlockItemList[43],
			0x106,
			0x01,
			0x249F0,
		},
		UnlockItem{
			UnlockItemList[44],
			0x119,
			0x01,
			0xEA60,
		},
		UnlockItem{
			UnlockItemList[45],
			0x11A,
			0x01,
			0x1D4C0,
		},
		UnlockItem{
			UnlockItemList[46],
			0x138,
			0x01,
			0x249F0,
		},
		UnlockItem{
			UnlockItemList[47],
			0x15C,
			0x01,
			0x7A120,
		},
		UnlockItem{
			UnlockItemList[48],
			0x182,
			0x01,
			0x186A0,
		},
		UnlockItem{
			UnlockItemList[49],
			0x183,
			0x01,
			0x186A0,
		},
		UnlockItem{
			UnlockItemList[50],
			0x184,
			0x01,
			0x186A0,
		},
		UnlockItem{
			UnlockItemList[51],
			0x1FA,
			0x01,
			0x7530,
		},
		UnlockItem{
			UnlockItemList[52],
			0x1FB,
			0x01,
			0xC350,
		},
		UnlockItem{
			UnlockItemList[53],
			0x1FC,
			0x01,
			0x30D40,
		},
		UnlockItem{
			UnlockItemList[54],
			0x207,
			0x00,
			0xA0,
		},
		UnlockItem{
			UnlockItemList[55],
			0x208,
			0x00,
			0x104,
		},
		UnlockItem{
			UnlockItemList[56],
			0x209,
			0x00,
			0x1E0,
		},
		UnlockItem{
			UnlockItemList[57],
			0x20A,
			0x00,
			0x244,
		},
		UnlockItem{
			UnlockItemList[58],
			0x258,
			0x00,
			0x244,
		},
		UnlockItem{
			UnlockItemList[59],
			0x259,
			0x00,
			0x30C,
		},
		UnlockItem{
			UnlockItemList[60],
			0x291,
			0x01,
			0x249F0,
		},
		UnlockItem{
			UnlockItemList[61],
			0x292,
			0x01,
			0x75300,
		},
		UnlockItem{
			UnlockItemList[62],
			0x293,
			0x01,
			0x35B60,
		},
		UnlockItem{
			UnlockItemList[63],
			0x294,
			0x00,
			0x140,
		},
		UnlockItem{
			UnlockItemList[64],
			0x295,
			0x00,
			0x208,
		},
		UnlockItem{
			UnlockItemList[65],
			0x31F,
			0x00,
			0x208,
		},
		UnlockItem{
			UnlockItemList[66],
			0x3A4,
			0x01,
			0x493E0,
		},
		UnlockItem{
			UnlockItemList[67],
			0x444,
			0x01,
			0x3A980,
		},
		UnlockItem{
			UnlockItemList[68],
			0x445,
			0x01,
			0x57E40,
		},
		UnlockItem{
			UnlockItemList[69],
			0x446,
			0x01,
			0x75300,
		},
		UnlockItem{
			UnlockItemList[70],
			0x4A9,
			0x01,
			0x249F0,
		},
		UnlockItem{
			UnlockItemList[71],
			0x4AA,
			0x01,
			0x30D40,
		},
		UnlockItem{
			UnlockItemList[72],
			0x4FC,
			0x01,
			0x9942,
		},
		UnlockItem{
			UnlockItemList[73],
			0x4FD,
			0x01,
			0x22986,
		},
		UnlockItem{
			UnlockItemList[74],
			0x4FE,
			0x01,
			0x2ED8C,
		},
	}
)
