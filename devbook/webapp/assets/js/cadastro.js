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
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        alert("Usuário cadastrado com sucesso!");
    }).fail(function (erro){
        console.log(erro)
        alert("Erro ao cadastrar o Usuário!")
    });
}