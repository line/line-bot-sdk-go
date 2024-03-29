/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py
package messaging_api

// AreaDemographic type
type AreaDemographic string

// AreaDemographic constants
const (

	// jp_01: 北海道 // Hokkaido

	AreaDemographic_HOKKAIDO AreaDemographic = "jp_01"

	// jp_02: 青森県 // Aomori

	AreaDemographic_AOMORI AreaDemographic = "jp_02"

	// jp_03: 岩手県 // Iwate

	AreaDemographic_IWATE AreaDemographic = "jp_03"

	// jp_04: 宮城県 // Miyagi

	AreaDemographic_MIYAGI AreaDemographic = "jp_04"

	// jp_05: 秋田県 // Akita

	AreaDemographic_AKITA AreaDemographic = "jp_05"

	// jp_06: 山形県 // Yamagata

	AreaDemographic_YAMAGATA AreaDemographic = "jp_06"

	// jp_07: 福島県 // Fukushima

	AreaDemographic_FUKUSHIMA AreaDemographic = "jp_07"

	// jp_08: 茨城県 // Ibaraki

	AreaDemographic_IBARAKI AreaDemographic = "jp_08"

	// jp_09: 栃木県 // Tochigi

	AreaDemographic_TOCHIGI AreaDemographic = "jp_09"

	// jp_10: 群馬県 // Gunma

	AreaDemographic_GUNMA AreaDemographic = "jp_10"

	// jp_11: 埼玉県 // Saitama

	AreaDemographic_SAITAMA AreaDemographic = "jp_11"

	// jp_12: 千葉県 // Chiba

	AreaDemographic_CHIBA AreaDemographic = "jp_12"

	// jp_13: 東京都 // Tokyo

	AreaDemographic_TOKYO AreaDemographic = "jp_13"

	// jp_14: 神奈川県 // Kanagawa

	AreaDemographic_KANAGAWA AreaDemographic = "jp_14"

	// jp_15: 新潟県 // Niigata

	AreaDemographic_NIIGATA AreaDemographic = "jp_15"

	// jp_16: 富山県 // Toyama

	AreaDemographic_TOYAMA AreaDemographic = "jp_16"

	// jp_17: 石川県 // Ishikawa

	AreaDemographic_ISHIKAWA AreaDemographic = "jp_17"

	// jp_18: 福井県 // Fukui

	AreaDemographic_FUKUI AreaDemographic = "jp_18"

	// jp_19: 山梨県 // Yamanashi

	AreaDemographic_YAMANASHI AreaDemographic = "jp_19"

	// jp_20: 長野県 // Nagano

	AreaDemographic_NAGANO AreaDemographic = "jp_20"

	// jp_21: 岐阜県 // Gifu

	AreaDemographic_GIFU AreaDemographic = "jp_21"

	// jp_22: 静岡県 // Shizuoka

	AreaDemographic_SHIZUOKA AreaDemographic = "jp_22"

	// jp_23: 愛知県 // Aichi

	AreaDemographic_AICHI AreaDemographic = "jp_23"

	// jp_24: 三重県 // Mie

	AreaDemographic_MIE AreaDemographic = "jp_24"

	// jp_25: 滋賀県 // Shiga

	AreaDemographic_SHIGA AreaDemographic = "jp_25"

	// jp_26: 京都府 // Kyoto

	AreaDemographic_KYOTO AreaDemographic = "jp_26"

	// jp_27: 大阪府 // Osaka

	AreaDemographic_OSAKA AreaDemographic = "jp_27"

	// jp_28: 兵庫県 // Hyougo

	AreaDemographic_HYOUGO AreaDemographic = "jp_28"

	// jp_29: 奈良県 // Nara

	AreaDemographic_NARA AreaDemographic = "jp_29"

	// jp_30: 和歌山県 // Wakayama

	AreaDemographic_WAKAYAMA AreaDemographic = "jp_30"

	// jp_31: 鳥取県 // Tottori

	AreaDemographic_TOTTORI AreaDemographic = "jp_31"

	// jp_32: 島根県 // Shimane

	AreaDemographic_SHIMANE AreaDemographic = "jp_32"

	// jp_33: 岡山県 // Okayama

	AreaDemographic_OKAYAMA AreaDemographic = "jp_33"

	// jp_34: 広島県 // Hiroshima

	AreaDemographic_HIROSHIMA AreaDemographic = "jp_34"

	// jp_35: 山口県 // Yamaguchi

	AreaDemographic_YAMAGUCHI AreaDemographic = "jp_35"

	// jp_36: 徳島県 // Tokushima

	AreaDemographic_TOKUSHIMA AreaDemographic = "jp_36"

	// jp_37: 香川県 // Kagawa

	AreaDemographic_KAGAWA AreaDemographic = "jp_37"

	// jp_38: 愛媛県 // Ehime

	AreaDemographic_EHIME AreaDemographic = "jp_38"

	// jp_39: 高知県 // Kouchi

	AreaDemographic_KOUCHI AreaDemographic = "jp_39"

	// jp_40: 福岡県 // Fukuoka

	AreaDemographic_FUKUOKA AreaDemographic = "jp_40"

	// jp_41: 佐賀県 // Saga

	AreaDemographic_SAGA AreaDemographic = "jp_41"

	// jp_42: 長崎県 // Nagasaki

	AreaDemographic_NAGASAKI AreaDemographic = "jp_42"

	// jp_43: 熊本県 // Kumamoto

	AreaDemographic_KUMAMOTO AreaDemographic = "jp_43"

	// jp_44: 大分県 // Oita

	AreaDemographic_OITA AreaDemographic = "jp_44"

	// jp_45: 宮崎県 // Miyazaki

	AreaDemographic_MIYAZAKI AreaDemographic = "jp_45"

	// jp_46: 鹿児島県 // Kagoshima

	AreaDemographic_KAGOSHIMA AreaDemographic = "jp_46"

	// jp_47: 沖縄県 // Okinawa

	AreaDemographic_OKINAWA AreaDemographic = "jp_47"

	// tw_01: 台北市 // Taipei City

	AreaDemographic_TAIPEI_CITY AreaDemographic = "tw_01"

	// tw_02: 新北市 // New Taipei City

	AreaDemographic_NEW_TAIPEI_CITY AreaDemographic = "tw_02"

	// tw_03: 桃園市 // Taoyuan City

	AreaDemographic_TAOYUAN_CITY AreaDemographic = "tw_03"

	// tw_04: 台中市 // Taichung City

	AreaDemographic_TAICHUNG_CITY AreaDemographic = "tw_04"

	// tw_05: 台南市 // Tainan City

	AreaDemographic_TAINAN_CITY AreaDemographic = "tw_05"

	// tw_06: 高雄市 // Kaohsiung City

	AreaDemographic_KAOHSIUNG_CITY AreaDemographic = "tw_06"

	// tw_07: 基隆市 // Keelung City

	AreaDemographic_KEELUNG_CITY AreaDemographic = "tw_07"

	// tw_08: 新竹市 // Hsinchu City

	AreaDemographic_HSINCHU_CITY AreaDemographic = "tw_08"

	// tw_09: 嘉義市 // Chiayi City

	AreaDemographic_CHIAYI_CITY AreaDemographic = "tw_09"

	// tw_10: 新竹県 // Hsinchu County

	AreaDemographic_HSINCHU_COUNTY AreaDemographic = "tw_10"

	// tw_11: 苗栗県 // Miaoli County

	AreaDemographic_MIAOLI_COUNTY AreaDemographic = "tw_11"

	// tw_12: 彰化県 // Changhua County

	AreaDemographic_CHANGHUA_COUNTY AreaDemographic = "tw_12"

	// tw_13: 南投県 // Nantou County

	AreaDemographic_NANTOU_COUNTY AreaDemographic = "tw_13"

	// tw_14: 雲林県 // Yunlin County

	AreaDemographic_YUNLIN_COUNTY AreaDemographic = "tw_14"

	// tw_15: 嘉義県 // Chiayi County

	AreaDemographic_CHIAYI_COUNTY AreaDemographic = "tw_15"

	// tw_16: 屏東県 // Pingtung County

	AreaDemographic_PINGTUNG_COUNTY AreaDemographic = "tw_16"

	// tw_17: 宜蘭県 // Yilan County

	AreaDemographic_YILAN_COUNTY AreaDemographic = "tw_17"

	// tw_18: 花蓮県 // Hualien County

	AreaDemographic_HUALIEN_COUNTY AreaDemographic = "tw_18"

	// tw_19: 台東県 // Taitung County

	AreaDemographic_TAITUNG_COUNTY AreaDemographic = "tw_19"

	// tw_20: 澎湖県 // Penghu County

	AreaDemographic_PENGHU_COUNTY AreaDemographic = "tw_20"

	// tw_21: 金門県 // Kinmen County

	AreaDemographic_KINMEN_COUNTY AreaDemographic = "tw_21"

	// tw_22: 連江県 // Lienchiang County

	AreaDemographic_LIENCHIANG_COUNTY AreaDemographic = "tw_22"

	// th_01: バンコク // Bangkok

	AreaDemographic_BANGKOK AreaDemographic = "th_01"

	// th_02: パタヤ // Pattaya

	AreaDemographic_PATTAYA AreaDemographic = "th_02"

	// th_03: 北部 // Northern

	AreaDemographic_NORTHERN AreaDemographic = "th_03"

	// th_04: 中央部 // Central

	AreaDemographic_CENTRAL AreaDemographic = "th_04"

	// th_05: 南部 // Southern

	AreaDemographic_SOUTHERN AreaDemographic = "th_05"

	// th_06: 東部 // Eastern

	AreaDemographic_EASTERN AreaDemographic = "th_06"

	// th_07: 東北部 // NorthEastern

	AreaDemographic_NORTHEASTERN AreaDemographic = "th_07"

	// th_08: 西部 // Western

	AreaDemographic_WESTERN AreaDemographic = "th_08"

	// id_01: バリ // Bali

	AreaDemographic_BALI AreaDemographic = "id_01"

	// id_02: バンドン // Bandung

	AreaDemographic_BANDUNG AreaDemographic = "id_02"

	// id_03: バンジャルマシン // Banjarmasin

	AreaDemographic_BANJARMASIN AreaDemographic = "id_03"

	// id_04: ジャボデタベック (ジャカルタ首都圏) // Jabodetabek

	AreaDemographic_JABODETABEK AreaDemographic = "id_04"

	// id_06: マカッサル // Makassar

	AreaDemographic_MAKASSAR AreaDemographic = "id_05"

	// id_07: メダン // Medan

	AreaDemographic_MEDAN AreaDemographic = "id_06"

	// id_08: パレンバン // Palembang

	AreaDemographic_PALEMBANG AreaDemographic = "id_07"

	// id_09: サマリンダ // Samarinda

	AreaDemographic_SAMARINDA AreaDemographic = "id_08"

	// id_10: スマラン // Semarang

	AreaDemographic_SEMARANG AreaDemographic = "id_09"

	// id_11: スラバヤ // Surabaya

	AreaDemographic_SURABAYA AreaDemographic = "id_10"

	// id_12: ジョグジャカルタ // Yogyakarta

	AreaDemographic_YOGYAKARTA AreaDemographic = "id_11"

	// id_05: その他のエリア // Lainnya

	AreaDemographic_LAINNYA AreaDemographic = "id_12"
)
