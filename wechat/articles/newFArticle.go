package article

const ()

//Artice 单图文素材
//参数	是否必须	说明
//title	是	标题
//thumb_media_id	是	图文消息的封面图片素材id（必须是永久mediaID）
//author	是	作者
//digest	是	图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
//show_cover_pic	是	是否显示封面，0为false，即不显示，1为true，即显示
//content	是	图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
//content_source_url	是	图文消息的原文地址，即点击“阅读原文”后的URL
type Artice struct {
	Title            string `json:"title"`
	ThumbMediaID     string `json:"thumb_media_id"`
	Author           string `json:"author"`
	Digest           string `json:"digest"`
	ShowCoverPic     string `json:"show_cover_pic"`
	Content          string `json:"content"`
	ContentSourceURL string `json:"content_source_url"`
}

//Artices 多图文素材
type Artices struct {
	ArticeNews []*Artice `json:"articles"`
}
