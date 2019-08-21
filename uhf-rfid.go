<<<<<<< HEAD
package UHFRFID

const(
	UserMem	= 0x00
	PwdMem 	= 0x01
	EPCMem 	= 0x02
	TIDMem 	= 0x03
)

type Memory struct {
	Len 	uint8 	// It specifies the number of 16-bit words to be read. The value is less then 120, can not be 0. Otherwise, it returns the parameters error message.
	Data 	[]uint8	// Be written words. The most significant byte of each word is first. Wdt specifies the array of the word to be written. For example, WordPtr equal 0x02, then the first word in Data write in the address 0x02 of designated Mem, the second word write in 0x03, etc.
}

type Tag struct {
	Pwd 	Memory // shall contain the kill and and/or access passwords, if passwords are implemented on the Tag. The kill password shall be stored at memory addresses 00h to 1Fh; the access password shall be stored at memory addresses 20h to 3Fh.
	EPC 	Memory // shall contain a Stored CRC at memory addresses 00h to 0Fh, a Stored PC at addresses 10h to 1Fh, a code (such as an EPC, and hereafter referred to as an EPC) that identifies the object to which the Tag is or AXCEZE ONE SERIES 800 UHF RFID Reader User's Manual V2.0 13 will be attached beginning at address 20h, and if the Tag implements Extended Protocol Control (XPC) then either one or two XPC word(s) beginning at address 210h.
	TID 	Memory // shall contain an 8-bit ISO/IEC 15963 allocation class identifier at memory locations 00h to 07h. TID memory shall contain sufficient identifying information above 07h for an Interrogator to uniquely identify the custom commands and/or optional features that a Tag supports. 
	User 	Memory // is optional. This area of different manufacturers is different. There is no user area in G2 tag of Inpinj Company. There are 28 words in Philips Company. Can write protect in four distinct banks. It means this memory is never writeable or not writeable under the non-safe state; only password area can set unreadable. 
}

type Mask struct{
	Adr 	uint8 	// It specifies the starting byte address for the memory mask. For example, MaskAdr = 0x00 specifies the first EPC bytes, MaskAdr = 0x01 specifies the second EPC bytes, etc.
	Len		uint8 	// It is the mask length. That a Tag compares against the memory location that begins at MaskAdr and ends MaskLen bytes later. MaskAdr + MaskLen must be less the length of ECP number. Otherwise, it returns the parameters error message.
}

/*specifies place*/
type Spec struct{
	Adr 	uint8 	// It specifies the starting word address for the memory read. For example, WordPtr = 00h specifies the first 16-bit memory word, WordPtr = 01h specifies the second 16-bit memory word, etc.
	Name	uint8 	// It specifies whether the Read accesses Password, EPC, TID, or User memory. 0x00: Password memory; 0x01: EPC memory; 0x02; TID memory; 0x03: User memory. Other values reserved. Other value when error occurred. 
	WordPtr uint8 	// One byte. It specifies the starting word address for the memory write. For example, WordPtr = 00h specifies the first 16-bit memory word, WordPtr = 01h specifies the second 16-bit memory word, etc.
	Mem 	Memory
}

/*Select: One byte, defined as follows*/
const(
	SltKill	= 0x00 // Control Kill Password protection setting.
	SltPwd 	= 0x01 // Control Access password protection setting.
	SltEPC 	= 0x02 // Control EPC memory protection setting.
	SltTID 	= 0x03 // Control TID memory protection setting.
	SltUser = 0x04 // Control User memory protection setting.
)

/*
SetProtect:
	When Select is 0x00 or 0x01, SetProtect means as follows
		readable and writeable from any state.
		permanently readable and writeable.
		readable and writeable from the secured state.
		never readable and writeable
	When Select is 0x02, 0x03 or 0x04, SetProtect means as follows
		writeable from any state.
		permanently writeable.
		writeable from the secured state.
		never writeable.
*/
const(
	AnyState 	= 0x00 
	Permanently = 0x01 
	SecuredState= 0x02 
	Never 		= 0x03 
)

type Protect struct{
	Select 	 	uint8 	// defined as follows
	SetProtect 	uint8 	
}
/*Condition: The condition of detecting tags.*/
const(
	Equal	= 0x00 // equal condition.
	Unequal	= 0x01 // unequal condition.
	Greater = 0x02 // greater than condition.
	Lower 	= 0x03 // lower than condition
)

type UID []uint8 // 8 bytes, it is 6B tag’s UID. The low byte is fist.

type Reader struct {
	Version 	[]uint8 // The first byte is version number; the second byte is sub-version number.	
	Type 		uint8 	// The reader type byte. 0x09 lines on AXCEZE ONE SERIES 800
	Tr_Type 	uint8 	// supported protocol information. Bit1 is 1 for18000-6C protocol; Bit0 is 1 for 18000-6B protocol.
	/*Frequency
		MaxFre(Bit7) | MaxFre(Bit6) | MinFre(Bit7) | MinFre(Bit6) | FreqBand
		0 0 0 0 User band
		0 0 0 1 Chinese band2
		0 0 1 0 US band
		0 0 1 1 Korean band
		0 1 0 0 RFU
		0 1 0 1 RFU
		… … … … …
		1 1 1 1 RFU

		User band: Fs = 902.6 + N * 0.4 (MHz), N∈ [0, 62].
		Chinese band2: Fs = 920.125 + N * 0.25 (MHz), N∈ [0, 19].
		US band: Fs = 902.75 + N * 0.5 (MHz), N∈ [0, 49].
		Korean band: Fs = 917.1 + N * 0.2 (MHz), N∈ [0, 31].
	*/
	MaxFre 		uint8 	// Bit7-Bit6 indicates Frequency Band and Bit5-Bit0 indicates the reader current maximum frequency.
	MinFre 		uint8 	// Bit7-Bit6 indicates Frequency Band and Bit5-Bit0 indicates the reader current minimum frequency (maximum frequency >= minimum frequency).
	Adr 		uint8 	// 
	ScanTime 	uint8 	// Inventory Scan Time. The default value is 0x0A (corresponding to 10*100ms=1s). The value range is 0x03~0xFF (corresponding to 3*100ms~255*100ms). When the host tries to set value 0x00~0x02 to InventoryScanTime, the reader will set it to 0x0A automatically. In various environments, the actual inventory scan time may be 0~75ms longer than the InventoryScanTime defined.
	BaudRate 	uint8 	// The serial port baud rate default value is 57600 bps. Defined as follows: 0x00 9600 bps, 0x01 19200 bps, 0x02 38400 bps, 0x05 57600 bps, 0x06 115200 bps
	Pwr 		uint8 	// New power. The default value is 30(about 30dBm), it range is 0~30.
	//Acousto-optic
	ActiveT	 	uint8 	// LED flash and buzzer tweet time. (ActiveT*50ms), the default value is 0. 0<=ActiveT<=255.
	SilentT	 	uint8 	// The LED and the buzzer silent time (SilentT *50ms), the default value is0. 0<= SilentT <=255.
	Times	 	uint8 	// LED flash and buzzer tweet times (0<=Times<=255), the default value is0
	/*Wiegand
		Parameter Connect:
			Wg_mode: Bit0: Select Wiegand format interface.
				=0 Wiegand 26bits format interface.
				=1 Wiegand 34bits format interface.
			Bit1: High-bit first or Low-bit first.
				=0 High-bit first.
				=1 Low-bit first.
			Bit2~Bit7: RFU. Default value is zero.
			Wg_Data_Inteval: Sending Data Delay (0 ~255)*10ms, the default value is 30.
			Wg_Pulse_Width: Data pulse width (1 ~255)*10us, the default value is 10.
			Wg_Pulse_Inteval: Data pulse interval width (1 ~255)*100us, the default value is 15.
	*/
	Wg_mode	 			uint8
	Wg_Data_Inteval	 	uint8
	Wg_Pulse_Width	 	uint8
	Wg_Pulse_Inteval	uint8
	/*WorkMode
		Byte1 Read_mode
		Byte2 Mode_state
		Byte3 Mem_Inven
		Byte4 First_Adr
		Byte5 Word_Num
		Byte6 Tag_Time
		
		Parameter Connect:
			Read_mode:
				Bit1 Bit0 Work Mode
				0    0 	  Answer Mode
				0    1 	  Scan Mode
				1    0 	  Trigger Mode(Low)
				1    1 	  Trigger Mode(High)
				Bit2~Bit7: RFU. Default value is zero.
				Notes: Answer mode, the following parameter is invalid.
			Mode_state: 
				Bit0: Protocol bit.
					=0 the reader support 18000-6C protocol.
					=1 the reader support 18000-6B protocol.
				Bit1: Output mode bit.
					=0 Wiegand output.
					=1 RS232/RS485 output.
				Bit2: Beep Enable.
					=0 on
					=1 off
				Bit3: Wiegand output, 18000-6C protocol. First_Adr is byte address or word address.
					=0 word address.
			 		=1 bytes address.
				Bit4: Syris485 Enable. It is invalid when Bit1 is zero.
					=0 Common 485
					=1 Syris 485
					When Bit4 = 1:
					Validity: 
						18000-6C protocol: Read accesses Password, EPC, TID, User memory, Inventory Single.
						18000-6B protocol: validity.
				Bit5~Bit7: RFU. Default value is zero.
			Mem_Inven: 
				It is valid when the reader supports 18000-6C protocol. 
				It specifies whether the Read accesses Password, EPC, TID, User memory, Inventory multiple, Inventory Single, EAS Alarm. 
					0x00: Password memory;
					0x01: EPC memory; 
					0x02; TID memory; 
					0x03: User memory; 
					0x04 Inventory multiple;
					0x05 Inventory Single;
					0x06: EAS Alarm. Otherwise, it returns the parameters error message.
			First_Adr: It specifies the starting data address for the memory read.
				Support 18000-6C: First_Adr = 0x00 specifies the first 16-bit memory word, First_Adr = 0x01 specifies the second 16-bit memory word, etc.
				Support 18000-6B: First_Adr = 0x00 specifies the first 8-bit memory byte, First_Adr = 0x01 specifies the second 8-bit memory byte, etc.
			Word_Num: Only RS232 RS232/RS485 output, it is valid. It specifies the number of word for the memory read. The value range is 1~32. Syris 485 Mode, the value range is 1~4.
			Tag_Time: Read Single Tag Delay (0 ~255)*1s. The default value is zero.
			Validity: 
				18000-6C protocol: Read accesses Password, EPC, TID, User memory, Inventory Single.
				18000-6B protocol: validity.
			Output Format Connect In The Scan Mode Or Trigger Mode:
				RS232/RS485, serial output format is as follows:
				Notes: RS232/RS485 serial output mode, these must be no tag in the effective field when set reader parameter.
	*/
	WorkMode 	[]uint8 // 6Bytes 
	
	Accuracy 	uint8 // EAS Alarm Accuracy. The default value is 8, it range is 0~8.
	OffsetTime 	uint8 // Syris485 response offset time (0 ~100)*1ms, the default value is 0.
	TriggerTime uint8 // Trigger offset time (0 ~254)*1s, the default value is 0. When TriggerTime is 255, means get the current trigger offset time.
}
=======
package UHFRFID

const(
	UserMem	= 0x00
	PwdMem 	= 0x01
	EPCMem 	= 0x02
	TIDMem 	= 0x03
)

type Memory struct {
	Len 	uint8 	// It specifies the number of 16-bit words to be read. The value is less then 120, can not be 0. Otherwise, it returns the parameters error message.
	Data 	[]uint8	// Be written words. The most significant byte of each word is first. Wdt specifies the array of the word to be written. For example, WordPtr equal 0x02, then the first word in Data write in the address 0x02 of designated Mem, the second word write in 0x03, etc.
}

type Tag struct {
	Pwd 	Memory // shall contain the kill and and/or access passwords, if passwords are implemented on the Tag. The kill password shall be stored at memory addresses 00h to 1Fh; the access password shall be stored at memory addresses 20h to 3Fh.
	EPC 	Memory // shall contain a Stored CRC at memory addresses 00h to 0Fh, a Stored PC at addresses 10h to 1Fh, a code (such as an EPC, and hereafter referred to as an EPC) that identifies the object to which the Tag is or AXCEZE ONE SERIES 800 UHF RFID Reader User's Manual V2.0 13 will be attached beginning at address 20h, and if the Tag implements Extended Protocol Control (XPC) then either one or two XPC word(s) beginning at address 210h.
	TID 	Memory // shall contain an 8-bit ISO/IEC 15963 allocation class identifier at memory locations 00h to 07h. TID memory shall contain sufficient identifying information above 07h for an Interrogator to uniquely identify the custom commands and/or optional features that a Tag supports. 
	User 	Memory // is optional. This area of different manufacturers is different. There is no user area in G2 tag of Inpinj Company. There are 28 words in Philips Company. Can write protect in four distinct banks. It means this memory is never writeable or not writeable under the non-safe state; only password area can set unreadable. 
}

type Mask struct{
	Adr 	uint8 	// It specifies the starting byte address for the memory mask. For example, MaskAdr = 0x00 specifies the first EPC bytes, MaskAdr = 0x01 specifies the second EPC bytes, etc.
	Len		uint8 	// It is the mask length. That a Tag compares against the memory location that begins at MaskAdr and ends MaskLen bytes later. MaskAdr + MaskLen must be less the length of ECP number. Otherwise, it returns the parameters error message.
}

/*specifies place*/
type Spec struct{
	Adr 	uint8 	// It specifies the starting word address for the memory read. For example, WordPtr = 00h specifies the first 16-bit memory word, WordPtr = 01h specifies the second 16-bit memory word, etc.
	Name	uint8 	// It specifies whether the Read accesses Password, EPC, TID, or User memory. 0x00: Password memory; 0x01: EPC memory; 0x02; TID memory; 0x03: User memory. Other values reserved. Other value when error occurred. 
	WordPtr uint8 	// One byte. It specifies the starting word address for the memory write. For example, WordPtr = 00h specifies the first 16-bit memory word, WordPtr = 01h specifies the second 16-bit memory word, etc.
	Mem 	Memory
}

/*Select: One byte, defined as follows*/
const(
	SltKill	= 0x00 // Control Kill Password protection setting.
	SltPwd 	= 0x01 // Control Access password protection setting.
	SltEPC 	= 0x02 // Control EPC memory protection setting.
	SltTID 	= 0x03 // Control TID memory protection setting.
	SltUser = 0x04 // Control User memory protection setting.
)

/*
SetProtect:
	When Select is 0x00 or 0x01, SetProtect means as follows
		readable and writeable from any state.
		permanently readable and writeable.
		readable and writeable from the secured state.
		never readable and writeable
	When Select is 0x02, 0x03 or 0x04, SetProtect means as follows
		writeable from any state.
		permanently writeable.
		writeable from the secured state.
		never writeable.
*/
const(
	AnyState 	= 0x00 
	Permanently = 0x01 
	SecuredState= 0x02 
	Never 		= 0x03 
)

type Protect struct{
	Select 	 	uint8 	// defined as follows
	SetProtect 	uint8 	
}
/*Condition: The condition of detecting tags.*/
const(
	Equal	= 0x00 // equal condition.
	Unequal	= 0x01 // unequal condition.
	Greater = 0x02 // greater than condition.
	Lower 	= 0x03 // lower than condition
)

type UID []uint8 // 8 bytes, it is 6B tag’s UID. The low byte is fist.

type Reader struct {
	Version 	[]uint8 // The first byte is version number; the second byte is sub-version number.	
	Type 		uint8 	// The reader type byte. 0x09 lines on AXCEZE ONE SERIES 800
	Tr_Type 	uint8 	// supported protocol information. Bit1 is 1 for18000-6C protocol; Bit0 is 1 for 18000-6B protocol.
	/*Frequency
		MaxFre(Bit7) | MaxFre(Bit6) | MinFre(Bit7) | MinFre(Bit6) | FreqBand
		0 0 0 0 User band
		0 0 0 1 Chinese band2
		0 0 1 0 US band
		0 0 1 1 Korean band
		0 1 0 0 RFU
		0 1 0 1 RFU
		… … … … …
		1 1 1 1 RFU

		User band: Fs = 902.6 + N * 0.4 (MHz), N∈ [0, 62].
		Chinese band2: Fs = 920.125 + N * 0.25 (MHz), N∈ [0, 19].
		US band: Fs = 902.75 + N * 0.5 (MHz), N∈ [0, 49].
		Korean band: Fs = 917.1 + N * 0.2 (MHz), N∈ [0, 31].
	*/
	MaxFre 		uint8 	// Bit7-Bit6 indicates Frequency Band and Bit5-Bit0 indicates the reader current maximum frequency.
	MinFre 		uint8 	// Bit7-Bit6 indicates Frequency Band and Bit5-Bit0 indicates the reader current minimum frequency (maximum frequency >= minimum frequency).
	Adr 		uint8 	// 
	ScanTime 	uint8 	// Inventory Scan Time. The default value is 0x0A (corresponding to 10*100ms=1s). The value range is 0x03~0xFF (corresponding to 3*100ms~255*100ms). When the host tries to set value 0x00~0x02 to InventoryScanTime, the reader will set it to 0x0A automatically. In various environments, the actual inventory scan time may be 0~75ms longer than the InventoryScanTime defined.
	BaudRate 	uint8 	// The serial port baud rate default value is 57600 bps. Defined as follows: 0x00 9600 bps, 0x01 19200 bps, 0x02 38400 bps, 0x05 57600 bps, 0x06 115200 bps
	Pwr 		uint8 	// New power. The default value is 30(about 30dBm), it range is 0~30.
	//Acousto-optic
	ActiveT	 	uint8 	// LED flash and buzzer tweet time. (ActiveT*50ms), the default value is 0. 0<=ActiveT<=255.
	SilentT	 	uint8 	// The LED and the buzzer silent time (SilentT *50ms), the default value is0. 0<= SilentT <=255.
	Times	 	uint8 	// LED flash and buzzer tweet times (0<=Times<=255), the default value is0
	/*Wiegand
		Parameter Connect:
			Wg_mode: Bit0: Select Wiegand format interface.
				=0 Wiegand 26bits format interface.
				=1 Wiegand 34bits format interface.
			Bit1: High-bit first or Low-bit first.
				=0 High-bit first.
				=1 Low-bit first.
			Bit2~Bit7: RFU. Default value is zero.
			Wg_Data_Inteval: Sending Data Delay (0 ~255)*10ms, the default value is 30.
			Wg_Pulse_Width: Data pulse width (1 ~255)*10us, the default value is 10.
			Wg_Pulse_Inteval: Data pulse interval width (1 ~255)*100us, the default value is 15.
	*/
	Wg_mode	 			uint8
	Wg_Data_Inteval	 	uint8
	Wg_Pulse_Width	 	uint8
	Wg_Pulse_Inteval	uint8
	/*WorkMode
		Byte1 Read_mode
		Byte2 Mode_state
		Byte3 Mem_Inven
		Byte4 First_Adr
		Byte5 Word_Num
		Byte6 Tag_Time
		
		Parameter Connect:
			Read_mode:
				Bit1 Bit0 Work Mode
				0    0 	  Answer Mode
				0    1 	  Scan Mode
				1    0 	  Trigger Mode(Low)
				1    1 	  Trigger Mode(High)
				Bit2~Bit7: RFU. Default value is zero.
				Notes: Answer mode, the following parameter is invalid.
			Mode_state: 
				Bit0: Protocol bit.
					=0 the reader support 18000-6C protocol.
					=1 the reader support 18000-6B protocol.
				Bit1: Output mode bit.
					=0 Wiegand output.
					=1 RS232/RS485 output.
				Bit2: Beep Enable.
					=0 on
					=1 off
				Bit3: Wiegand output, 18000-6C protocol. First_Adr is byte address or word address.
					=0 word address.
			 		=1 bytes address.
				Bit4: Syris485 Enable. It is invalid when Bit1 is zero.
					=0 Common 485
					=1 Syris 485
					When Bit4 = 1:
					Validity: 
						18000-6C protocol: Read accesses Password, EPC, TID, User memory, Inventory Single.
						18000-6B protocol: validity.
				Bit5~Bit7: RFU. Default value is zero.
			Mem_Inven: 
				It is valid when the reader supports 18000-6C protocol. 
				It specifies whether the Read accesses Password, EPC, TID, User memory, Inventory multiple, Inventory Single, EAS Alarm. 
					0x00: Password memory;
					0x01: EPC memory; 
					0x02; TID memory; 
					0x03: User memory; 
					0x04 Inventory multiple;
					0x05 Inventory Single;
					0x06: EAS Alarm. Otherwise, it returns the parameters error message.
			First_Adr: It specifies the starting data address for the memory read.
				Support 18000-6C: First_Adr = 0x00 specifies the first 16-bit memory word, First_Adr = 0x01 specifies the second 16-bit memory word, etc.
				Support 18000-6B: First_Adr = 0x00 specifies the first 8-bit memory byte, First_Adr = 0x01 specifies the second 8-bit memory byte, etc.
			Word_Num: Only RS232 RS232/RS485 output, it is valid. It specifies the number of word for the memory read. The value range is 1~32. Syris 485 Mode, the value range is 1~4.
			Tag_Time: Read Single Tag Delay (0 ~255)*1s. The default value is zero.
			Validity: 
				18000-6C protocol: Read accesses Password, EPC, TID, User memory, Inventory Single.
				18000-6B protocol: validity.
			Output Format Connect In The Scan Mode Or Trigger Mode:
				RS232/RS485, serial output format is as follows:
				Notes: RS232/RS485 serial output mode, these must be no tag in the effective field when set reader parameter.
	*/
	WorkMode 	[]uint8 // 6Bytes 
	
	Accuracy 	uint8 // EAS Alarm Accuracy. The default value is 8, it range is 0~8.
	OffsetTime 	uint8 // Syris485 response offset time (0 ~100)*1ms, the default value is 0.
	TriggerTime uint8 // Trigger offset time (0 ~254)*1s, the default value is 0. When TriggerTime is 255, means get the current trigger offset time.
}
>>>>>>> 8/22
