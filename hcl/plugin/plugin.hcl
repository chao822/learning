io_mode = "async"

plugin "plugin_type1" "plugin_name1" {
  plugin_cmd = "./plugin_type1/plugin_type1.so"
  plugin_string_data = "test1"
#  plugin_int_data = 2
#  plugin_block_data "plugin_block_type" "plugin_block_name" {
#    test = "test"
#  }
}

#provider "provider_type2" "provider_name2" {
#  plugin_cmd = "yyy"
#  plugin_string_data1 = "test2"
#}