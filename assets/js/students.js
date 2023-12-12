function onDelete(id) {
  let response = confirm("Tem certeza que você deseja deletar este aluno?")
  if (response) {
    window.location = "/delete?id=" + id
  }
}

function addStudent(lenStudents) {
  if (lenStudents >= 10) {
    alert("O número máximo de estudantes (10) já foi atingido!")
  } else {
    window.location = "/add-student"
  }
}