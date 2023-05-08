io_mode = "async"

plugin "plugin_type1" "plugin_name1" {
  plugin_cmd = "./plugin_type1/plugin_type1.so"
  plugin_string_data = "test1"
  plugin_int_data = 2
  plugin_block_data {
    test = "test"
  }
}

plugin "plugin_type2" "plugin_name2" {
  plugin_cmd = "./plugin_type2/plugin_type2.so"
  plugin_string_data = "test2"
}