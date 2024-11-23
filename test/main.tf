provider "pensando" {
  dsc_address = var.dsc_address 
  dsc_port    = var.dsc_port 
}

data "pensando_device_status" "dsc_status" {
}

output "dsc_status" {
  value = data.pensando_device_status.dsc_status
}
