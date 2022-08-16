# 将.png格式转为 .basis
basisu xxx.png
# 针对法线/金属/粗糙贴图等linear颜色空间的贴图 需加上-linear
basisu xxx.png -linear
# 最大限度保证图片质量的转换
basisu xxx.png -comp_level 5 -max_endpoints 16128 -max_selectors 16128 -no_selector_rdo
# 最大限度压缩linear颜色空间的贴图
basisu xxx.png -linear -global_sel_pal -no_hybrid_sel_cb