# Example_Go


 Kí hiệu: _ (blank identifier) : bỏ qua giá trị mà ko có nhu cầu sử dụng


Trong Go, nil là giá trị 0 cho con trỏ, giao diện, bản đồ, lát, kênh và loại chức năng, đại diện cho một chưa khởi tạo giá trị.

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

Từ khóa go gọi trước hàm hoặc phương thức, chúng ta sẽ có một Goroutine chạy đồng thời
