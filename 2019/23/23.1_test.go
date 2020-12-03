package _2019

import (
	"reflect"
	"testing"
)

func TestGetFirstPacketSentToNAT(t *testing.T) {
	tests := []struct {
		name         string
		instructions []int
		want         Packet
	}{
		{
			name:         "input for first part",
			instructions: []int{3, 62, 1001, 62, 11, 10, 109, 2229, 105, 1, 0, 1856, 918, 1151, 1392, 2070, 1460, 2099, 660, 1714, 1217, 691, 1584, 1755, 1489, 600, 2198, 571, 1617, 2165, 885, 1897, 1679, 1054, 1524, 1023, 1250, 1827, 821, 852, 1427, 1648, 990, 629, 1285, 1967, 955, 722, 2134, 1553, 1316, 1790, 1361, 2037, 1085, 1120, 1928, 1184, 786, 753, 1996, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3, 64, 1008, 64, -1, 62, 1006, 62, 88, 1006, 61, 170, 1105, 1, 73, 3, 65, 20101, 0, 64, 1, 20102, 1, 66, 2, 21101, 0, 105, 0, 1105, 1, 436, 1201, 1, -1, 64, 1007, 64, 0, 62, 1005, 62, 73, 7, 64, 67, 62, 1006, 62, 73, 1002, 64, 2, 133, 1, 133, 68, 133, 102, 1, 0, 62, 1001, 133, 1, 140, 8, 0, 65, 63, 2, 63, 62, 62, 1005, 62, 73, 1002, 64, 2, 161, 1, 161, 68, 161, 1101, 1, 0, 0, 1001, 161, 1, 169, 102, 1, 65, 0, 1102, 1, 1, 61, 1101, 0, 0, 63, 7, 63, 67, 62, 1006, 62, 203, 1002, 63, 2, 194, 1, 68, 194, 194, 1006, 0, 73, 1001, 63, 1, 63, 1106, 0, 178, 21101, 210, 0, 0, 105, 1, 69, 2101, 0, 1, 70, 1102, 1, 0, 63, 7, 63, 71, 62, 1006, 62, 250, 1002, 63, 2, 234, 1, 72, 234, 234, 4, 0, 101, 1, 234, 240, 4, 0, 4, 70, 1001, 63, 1, 63, 1106, 0, 218, 1105, 1, 73, 109, 4, 21101, 0, 0, -3, 21101, 0, 0, -2, 20207, -2, 67, -1, 1206, -1, 293, 1202, -2, 2, 283, 101, 1, 283, 283, 1, 68, 283, 283, 22001, 0, -3, -3, 21201, -2, 1, -2, 1105, 1, 263, 22102, 1, -3, -3, 109, -4, 2106, 0, 0, 109, 4, 21102, 1, 1, -3, 21102, 0, 1, -2, 20207, -2, 67, -1, 1206, -1, 342, 1202, -2, 2, 332, 101, 1, 332, 332, 1, 68, 332, 332, 22002, 0, -3, -3, 21201, -2, 1, -2, 1106, 0, 312, 21201, -3, 0, -3, 109, -4, 2106, 0, 0, 109, 1, 101, 1, 68, 358, 21001, 0, 0, 1, 101, 3, 68, 366, 21002, 0, 1, 2, 21101, 0, 376, 0, 1106, 0, 436, 22102, 1, 1, 0, 109, -1, 2106, 0, 0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536, 131072, 262144, 524288, 1048576, 2097152, 4194304, 8388608, 16777216, 33554432, 67108864, 134217728, 268435456, 536870912, 1073741824, 2147483648, 4294967296, 8589934592, 17179869184, 34359738368, 68719476736, 137438953472, 274877906944, 549755813888, 1099511627776, 2199023255552, 4398046511104, 8796093022208, 17592186044416, 35184372088832, 70368744177664, 140737488355328, 281474976710656, 562949953421312, 1125899906842624, 109, 8, 21202, -6, 10, -5, 22207, -7, -5, -5, 1205, -5, 521, 21102, 1, 0, -4, 21102, 0, 1, -3, 21101, 51, 0, -2, 21201, -2, -1, -2, 1201, -2, 385, 470, 21001, 0, 0, -1, 21202, -3, 2, -3, 22207, -7, -1, -5, 1205, -5, 496, 21201, -3, 1, -3, 22102, -1, -1, -5, 22201, -7, -5, -7, 22207, -3, -6, -5, 1205, -5, 515, 22102, -1, -6, -5, 22201, -3, -5, -3, 22201, -1, -4, -4, 1205, -2, 461, 1105, 1, 547, 21101, 0, -1, -4, 21202, -6, -1, -6, 21207, -7, 0, -5, 1205, -5, 547, 22201, -7, -6, -7, 21201, -4, 1, -4, 1106, 0, 529, 21202, -4, 1, -7, 109, -8, 2106, 0, 0, 109, 1, 101, 1, 68, 564, 20101, 0, 0, 0, 109, -1, 2106, 0, 0, 1101, 39581, 0, 66, 1102, 1, 1, 67, 1101, 598, 0, 68, 1101, 0, 556, 69, 1101, 0, 0, 71, 1102, 600, 1, 72, 1105, 1, 73, 1, 1145, 1101, 83477, 0, 66, 1102, 1, 1, 67, 1101, 0, 627, 68, 1101, 0, 556, 69, 1102, 0, 1, 71, 1101, 629, 0, 72, 1106, 0, 73, 1, 1537, 1102, 23167, 1, 66, 1101, 0, 1, 67, 1101, 656, 0, 68, 1101, 0, 556, 69, 1102, 1, 1, 71, 1101, 0, 658, 72, 1105, 1, 73, 1, -78, 13, 778, 1101, 46757, 0, 66, 1101, 1, 0, 67, 1102, 1, 687, 68, 1101, 0, 556, 69, 1101, 0, 1, 71, 1102, 1, 689, 72, 1106, 0, 73, 1, 125, 1, 37767, 1101, 4591, 0, 66, 1101, 0, 1, 67, 1101, 718, 0, 68, 1101, 556, 0, 69, 1102, 1, 1, 71, 1102, 720, 1, 72, 1106, 0, 73, 1, 107, 49, 200722, 1102, 1, 2579, 66, 1102, 1, 1, 67, 1102, 749, 1, 68, 1102, 1, 556, 69, 1101, 1, 0, 71, 1102, 1, 751, 72, 1106, 0, 73, 1, 1861, 47, 80326, 1102, 1, 55057, 66, 1102, 2, 1, 67, 1101, 0, 780, 68, 1101, 0, 302, 69, 1101, 1, 0, 71, 1101, 0, 784, 72, 1106, 0, 73, 0, 0, 0, 0, 31, 68881, 1101, 40163, 0, 66, 1102, 1, 3, 67, 1101, 0, 813, 68, 1101, 302, 0, 69, 1102, 1, 1, 71, 1102, 819, 1, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 35, 12577, 1102, 1, 10753, 66, 1102, 1, 1, 67, 1101, 0, 848, 68, 1102, 556, 1, 69, 1101, 0, 1, 71, 1101, 850, 0, 72, 1106, 0, 73, 1, 13, 49, 301083, 1101, 0, 69767, 66, 1102, 1, 1, 67, 1101, 0, 879, 68, 1102, 556, 1, 69, 1101, 2, 0, 71, 1102, 881, 1, 72, 1106, 0, 73, 1, 10, 1, 50356, 8, 4146, 1102, 3391, 1, 66, 1102, 2, 1, 67, 1101, 912, 0, 68, 1102, 351, 1, 69, 1101, 1, 0, 71, 1101, 916, 0, 72, 1105, 1, 73, 0, 0, 0, 0, 255, 51059, 1101, 12589, 0, 66, 1102, 4, 1, 67, 1101, 945, 0, 68, 1101, 0, 302, 69, 1102, 1, 1, 71, 1101, 953, 0, 72, 1106, 0, 73, 0, 0, 0, 0, 0, 0, 0, 0, 8, 3455, 1101, 0, 12577, 66, 1102, 1, 3, 67, 1101, 0, 982, 68, 1101, 253, 0, 69, 1102, 1, 1, 71, 1101, 0, 988, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 18, 96181, 1102, 1, 68881, 66, 1102, 2, 1, 67, 1102, 1, 1017, 68, 1101, 0, 302, 69, 1101, 0, 1, 71, 1101, 1021, 0, 72, 1106, 0, 73, 0, 0, 0, 0, 2, 9473, 1101, 10957, 0, 66, 1101, 0, 1, 67, 1101, 1050, 0, 68, 1101, 556, 0, 69, 1101, 0, 1, 71, 1101, 0, 1052, 72, 1106, 0, 73, 1, -87195, 6, 188913, 1101, 0, 48187, 66, 1101, 1, 0, 67, 1102, 1, 1081, 68, 1101, 556, 0, 69, 1101, 0, 1, 71, 1102, 1083, 1, 72, 1105, 1, 73, 1, 4, 49, 100361, 1102, 81563, 1, 66, 1101, 0, 3, 67, 1101, 1112, 0, 68, 1102, 302, 1, 69, 1101, 0, 1, 71, 1101, 1118, 0, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 35, 37731, 1102, 72053, 1, 66, 1102, 1, 1, 67, 1101, 0, 1147, 68, 1102, 1, 556, 69, 1101, 1, 0, 71, 1101, 0, 1149, 72, 1105, 1, 73, 1, -686, 47, 120489, 1102, 1, 9473, 66, 1101, 0, 2, 67, 1101, 0, 1178, 68, 1101, 0, 302, 69, 1102, 1, 1, 71, 1102, 1182, 1, 72, 1106, 0, 73, 0, 0, 0, 0, 11, 74357, 1101, 0, 19753, 66, 1102, 2, 1, 67, 1102, 1, 1211, 68, 1102, 302, 1, 69, 1102, 1, 1, 71, 1101, 0, 1215, 72, 1106, 0, 73, 0, 0, 0, 0, 40, 9754, 1101, 16979, 0, 66, 1102, 1, 1, 67, 1101, 1244, 0, 68, 1101, 556, 0, 69, 1101, 2, 0, 71, 1102, 1, 1246, 72, 1106, 0, 73, 1, 641, 49, 602166, 13, 389, 1101, 1291, 0, 66, 1102, 3, 1, 67, 1101, 1277, 0, 68, 1101, 0, 302, 69, 1102, 1, 1, 71, 1102, 1, 1283, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 40, 14631, 1101, 90073, 0, 66, 1102, 1, 1, 67, 1101, 1312, 0, 68, 1102, 556, 1, 69, 1102, 1, 1, 71, 1102, 1, 1314, 72, 1105, 1, 73, 1, 52895, 6, 125942, 1101, 22769, 0, 66, 1102, 1, 1, 67, 1102, 1, 1343, 68, 1102, 556, 1, 69, 1102, 1, 8, 71, 1102, 1345, 1, 72, 1105, 1, 73, 1, 1, 48, 55057, 31, 137762, 2, 18946, 11, 148714, 29, 97978, 43, 244689, 47, 40163, 13, 1167, 1101, 0, 1777, 66, 1101, 1, 0, 67, 1102, 1, 1388, 68, 1102, 556, 1, 69, 1102, 1, 1, 71, 1102, 1390, 1, 72, 1105, 1, 73, 1, 51973, 48, 110114, 1101, 28069, 0, 66, 1102, 3, 1, 67, 1102, 1, 1419, 68, 1102, 1, 302, 69, 1101, 0, 1, 71, 1101, 0, 1425, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 46, 39506, 1102, 48989, 1, 66, 1102, 2, 1, 67, 1101, 0, 1454, 68, 1102, 1, 302, 69, 1102, 1, 1, 71, 1102, 1458, 1, 72, 1106, 0, 73, 0, 0, 0, 0, 43, 81563, 1101, 0, 46861, 66, 1102, 1, 1, 67, 1101, 0, 1487, 68, 1101, 0, 556, 69, 1101, 0, 0, 71, 1102, 1, 1489, 72, 1105, 1, 73, 1, 1696, 1102, 1, 389, 66, 1101, 3, 0, 67, 1101, 1516, 0, 68, 1101, 0, 302, 69, 1101, 0, 1, 71, 1102, 1, 1522, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 21, 11093, 1102, 1, 15887, 66, 1101, 0, 1, 67, 1101, 1551, 0, 68, 1102, 556, 1, 69, 1102, 0, 1, 71, 1101, 0, 1553, 72, 1106, 0, 73, 1, 1653, 1102, 25693, 1, 66, 1102, 1, 1, 67, 1102, 1580, 1, 68, 1102, 1, 556, 69, 1101, 0, 1, 71, 1102, 1, 1582, 72, 1106, 0, 73, 1, 51, 43, 163126, 1101, 0, 74357, 66, 1102, 2, 1, 67, 1102, 1611, 1, 68, 1101, 0, 302, 69, 1101, 0, 1, 71, 1101, 1615, 0, 72, 1106, 0, 73, 0, 0, 0, 0, 49, 501805, 1102, 4597, 1, 66, 1101, 0, 1, 67, 1101, 1644, 0, 68, 1102, 556, 1, 69, 1101, 0, 1, 71, 1101, 0, 1646, 72, 1105, 1, 73, 1, 319931, 6, 62971, 1102, 1, 62467, 66, 1102, 1, 1, 67, 1102, 1675, 1, 68, 1102, 1, 556, 69, 1101, 0, 1, 71, 1101, 1677, 0, 72, 1105, 1, 73, 1, 41, 3, 28069, 1101, 11093, 0, 66, 1102, 3, 1, 67, 1102, 1706, 1, 68, 1102, 1, 302, 69, 1102, 1, 1, 71, 1102, 1, 1712, 72, 1106, 0, 73, 0, 0, 0, 0, 0, 0, 40, 4877, 1102, 1, 691, 66, 1102, 1, 6, 67, 1101, 0, 1741, 68, 1102, 302, 1, 69, 1102, 1, 1, 71, 1101, 1753, 0, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 19, 6782, 1102, 64567, 1, 66, 1101, 1, 0, 67, 1102, 1782, 1, 68, 1102, 1, 556, 69, 1101, 3, 0, 71, 1101, 0, 1784, 72, 1105, 1, 73, 1, 5, 1, 12589, 1, 25178, 8, 2764, 1101, 0, 4877, 66, 1102, 4, 1, 67, 1102, 1817, 1, 68, 1101, 0, 253, 69, 1102, 1, 1, 71, 1102, 1825, 1, 72, 1106, 0, 73, 0, 0, 0, 0, 0, 0, 0, 0, 19, 3391, 1101, 0, 29863, 66, 1102, 1, 1, 67, 1101, 0, 1854, 68, 1102, 556, 1, 69, 1102, 0, 1, 71, 1101, 0, 1856, 72, 1106, 0, 73, 1, 1846, 1101, 0, 51059, 66, 1101, 0, 1, 67, 1102, 1883, 1, 68, 1101, 556, 0, 69, 1102, 1, 6, 71, 1102, 1, 1885, 72, 1106, 0, 73, 1, 25793, 46, 19753, 21, 22186, 21, 33279, 25, 1291, 25, 2582, 25, 3873, 1101, 0, 75781, 66, 1101, 0, 1, 67, 1102, 1, 1924, 68, 1102, 1, 556, 69, 1101, 1, 0, 71, 1102, 1926, 1, 72, 1105, 1, 73, 1, 160, 8, 1382, 1101, 99469, 0, 66, 1102, 1, 1, 67, 1101, 0, 1955, 68, 1101, 556, 0, 69, 1101, 5, 0, 71, 1101, 0, 1957, 72, 1106, 0, 73, 1, 2, 49, 401444, 18, 192362, 3, 84207, 8, 691, 8, 2073, 1101, 0, 8761, 66, 1102, 1, 1, 67, 1101, 1994, 0, 68, 1102, 556, 1, 69, 1101, 0, 0, 71, 1101, 1996, 0, 72, 1105, 1, 73, 1, 1369, 1102, 100361, 1, 66, 1102, 1, 6, 67, 1102, 1, 2023, 68, 1101, 0, 302, 69, 1102, 1, 1, 71, 1101, 0, 2035, 72, 1106, 0, 73, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 40, 19508, 1102, 37967, 1, 66, 1101, 2, 0, 67, 1102, 1, 2064, 68, 1102, 1, 302, 69, 1102, 1, 1, 71, 1102, 2068, 1, 72, 1105, 1, 73, 0, 0, 0, 0, 35, 25154, 1101, 18367, 0, 66, 1101, 1, 0, 67, 1102, 2097, 1, 68, 1101, 0, 556, 69, 1102, 0, 1, 71, 1101, 2099, 0, 72, 1106, 0, 73, 1, 1503, 1101, 62971, 0, 66, 1102, 3, 1, 67, 1101, 2126, 0, 68, 1102, 1, 253, 69, 1101, 0, 1, 71, 1102, 1, 2132, 72, 1105, 1, 73, 0, 0, 0, 0, 0, 0, 42, 37967, 1102, 1, 37691, 66, 1101, 1, 0, 67, 1101, 2161, 0, 68, 1101, 0, 556, 69, 1102, 1, 1, 71, 1101, 0, 2163, 72, 1105, 1, 73, 1, 34, 42, 75934, 1102, 96181, 1, 66, 1102, 2, 1, 67, 1101, 2192, 0, 68, 1101, 0, 302, 69, 1101, 0, 1, 71, 1101, 0, 2196, 72, 1106, 0, 73, 0, 0, 0, 0, 3, 56138, 1102, 1, 44777, 66, 1101, 0, 1, 67, 1101, 2225, 0, 68, 1101, 556, 0, 69, 1102, 1, 1, 71, 1101, 2227, 0, 72, 1105, 1, 73, 1, 17929, 29, 48989},
			want:         Packet{x: 51059, y: 23815},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirstPacketSentToNAT(tt.instructions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFirstPacketSentToNAT() = %v, want %v", got, tt.want)
			}
		})
	}
}
