# removeConfig
Iskratel OLT Tool to remove running-config elements using a copy-paste stdin-stdout method

| No Flags, runs and waits for STDIN |

# Example Input
```sh
interface 0/8/1
onu serial-number ISKT019ECE4A
service-profile 101_CWMP_Uni
service-profile 102_DATA_Acc
exit

interface 0/8/2
onu serial-number ISKT861ADFF0
service-profile 101_CWMP
exit

interface 0/8/3
onu serial-number ISKT019ECE96
service-profile 101_CWMP_Uni
service-profile 102_DATA_Acc
exit



```

# Example Output
```sh
interface 0/8/1
no onu serial-number
no service-profile 101_CWMP_Uni
no service-profile 102_DATA_Acc
exit
interface 0/8/2
no onu serial-number
no service-profile 101_CWMP
exit
interface 0/8/3
no onu serial-number
no service-profile 101_CWMP_Uni
no service-profile 102_DATA_Acc
exit
```

# Copy-Paste Output to "config" menu on OLT
```sh
LBB-G16#config
LBB-G16(Config)#interface 0/8/1
LBB-G16(Interface 0/8/1)#no onu serial-number
LBB-G16(Interface 0/8/1)#no service-profile 101_CWMP_Uni
LBB-G16(Interface 0/8/1)#no service-profile 102_DATA_Acc
LBB-G16(Interface 0/8/1)#exit
LBB-G16(Config)#interface 0/8/2
LBB-G16(Interface 0/8/2)#no onu serial-number
LBB-G16(Interface 0/8/2)#no service-profile 101_CWMP
LBB-G16(Interface 0/8/2)#exit
LBB-G16(Config)#interface 0/8/3
LBB-G16(Interface 0/8/3)#no onu serial-number
LBB-G16(Interface 0/8/3)#no service-profile 101_CWMP_Uni
LBB-G16(Interface 0/8/3)#no service-profile 102_DATA_Acc
LBB-G16(Interface 0/8/3)#exit
LBB-G16(Config)#ex
LBB-G16#
```
