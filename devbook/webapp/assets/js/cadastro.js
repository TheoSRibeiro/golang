$('#formulario-cadastro').on('submit', criarUsuario);


function criarUsuario(evento){
    evento.preventDefault();
    console.log("Funcao criarUsuario!");
    
    //Comparar as senhas passadas
    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("As senhas não coicidem!")
        return;
    }

    //Requisicao AJAX - caso as senhas forem iguais
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            nome: $('#email').val(),
            nome: $('#nick').val(),
            nome: $('#senha').val()
        }
    });
}