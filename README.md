# Example_Go

- Các kiểu số nguyên trong Go là uint8, uint16, uint32, uint64, int8, int16, int32, int64. Các con số 8, 16, 32, 64 là số bít mà máy tính dùng để biểu diễn số nguyên đó. uint tức là unsigned int – là kiểu số nguyên không âm.
Các kiểu số thực trong Go là float32 và float64
Các kiểu số phức là complex64 và complex128

- Kí hiệu: _ (blank identifier) : bỏ qua giá trị mà ko có nhu cầu sử dụng

- Trong Go, ":=" toán tử là một phím tắt để khai báo và gởi tạo một biến trong một dòng (Go sử dụng giá trị ở bên phải để xác định kiểu của biến) 

- Trong Go, nil là giá trị 0 cho con trỏ, giao diện, bản đồ, lát, kênh và loại chức năng, đại diện cho một chưa khởi tạo giá trị.

nil không có nghĩa là một số trạng thái "không xác định", bản thân nó là một giá trị phù hợp. Một đối tượng trong Go là nil chỉ đơn giản là khi và chỉ khi giá trị của nó là nil, mà chỉ có thể là nếu nó thuộc một trong các loại đã nói ở trên.

error là một giao diện, vì vậy nil là một giá trị hợp lệ cho một, không giống như một string. Vì các lý do rõ ràng, lỗi nil đại diện cho không có lỗi.

nil trong Go chỉ đơn giản là giá trị con trỏ Null của các ngôn ngữ khác.

 error được tích hợp sẵn trong Go với giá trị mặc đinh là nil. 
 Có một cách xử lý lỗi là trả nó về như giá trị cuối cùng của hàm gọi và kiểm tra xem có nil hay không
VD: 
val, err := myFunction( args... );
if err != nil {
  // Xử lý lỗi
} else {
  // Thành công
}

- Từ khóa go gọi trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời
 
 - Hàm make() của golang tạo ra empty slice trống. Các empty sclice trống chứa tham chiếu mảng trống 
 Golang make () là một hàm Slice tích hợp được sử dụng để tạo một Slice. Hàm make () nhận ba đối số: kiểu, độ dài và dung lượng và trả về slice. Để tạo mảng có kích thước động, hãy sử dụng hàm make ().
  
- Map là một kiểu dữ liệu có sãn trong Go, một map là tập hợp các cặp key/value. Value chỉ được truy xuất bởi key tương ứng.
Một map được khởi tạo bằng cách truyền dữ liệu của key và value thông qua hàm "make".
Cú pháp make(map[type of key]type of value)


- Hàm Printf(<định dạng>, a) cho phép chúng ta in dữ liệu theo các dạng nhất định, tham số đầu tiên của hàm này là dạng chuỗi sẽ được in ra, các chuỗi này có kí tự đầu tiên là %, các kí tự tiếp theo mô tả kiểu in của dữ liệu, tham số thứ 2 là dữ liệu, trong ví dụ trên thì chúng ta có định nghĩa một struct tên là point có 2 trường kiểu int là x, y, tiếp theo chúng ta in một số dữ liệu ra màn hình bằng hàm Printf():

%v: in các giá trị của một biến struct
%+v: in các giá trị kèm với tên trường của biến struct
%#v: giống %+v kèm theo tên kiểu dữ liệu của struct và tên hàm đã gọi nó
%T: in tên struct và tên hàm đã gọi nó
%t: in giá trị boolean
%d: in số nguyên (hệ 10)
%b: in số nguyên dưới dạng số nhị phân (hệ 2)
%c: in kí tự dựa theo mã ASCII
%x: in số nguyên dưới dạng số thập lục phân (hệ 16) hoặc chuyển một chuỗi thành số thập lục phân
%f: in số thập phân
%e và %E: in số thập phân dưới dạng số mũ
%s: in một chuỗi
%q: in một chuỗi có 2 cặp dấu nháy kép “”
%6d: in một số nguyên, nếu số đó không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
%6.2f: in số thập phân, làm tròn đến 2 chữ số thập phân, sau đó nếu phần thập phân và phần nguyên cùng với dấu chấm không đủ 6 kí tự thì tự động thêm các khoảng trống vào bên trái cho đủ 6 kí tự
%-6.2f: tương tự với %6.2f nhưng các khoảng trống được thêm vào bên phải
%6s: in một chuỗi, nếu chuỗi không đủ 6 kí tự thì thêm các khoảng trống vào bên trái cho đủ
%-6s: tương tự %6s nhưng thêm các khoảng trống vào bên phải

Ngoài ra chúng ta có hàm Sprintf() chỉ trả về một chuỗi chứ không in ra màn hình.

Kiểu dữ liệu:
- bool : true
- string: "213"
- int: 213
- float64: 213.213
- var a byte / a = 2: uint8 


Tài liệu:

https://github.com/mmcgrana/gobyexample

https://github.com/callicoder/golang-tutorials

https://www.callicoder.com/

https://github.com/video-dev/hls.js/

